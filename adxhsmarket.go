package adxhsmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var errNonNilContext = errors.New("context must be non-nil")

const (
	// 小红书开放平台请求域名 api_domain
	baseURL = "https://adapi.xiaohongshu.com"
)

var (
	hdrAuthorizationKey = http.CanonicalHeaderKey("Access-Token")
	hdrContentTypeKey   = http.CanonicalHeaderKey("Content-Type")
	hdrUserAgentKey     = http.CanonicalHeaderKey("User-Agent")
)

var (
	contentTypeJSON = "application/json"
)

type Client struct {
	config *Config

	common service

	Auth       *AuthService
	Account    *AccountService
	Kol        *KolService
	Note       *NoteService
	Campaign   *CampaignService
	Unit       *UnitService
	Creativity *CreativityService
	Report     *ReportService
}

type service struct {
	client *Client
}

type Config struct {
	BaseURL          *url.URL
	AppId            int
	AppSecret        string
	HttpClient       *http.Client
	EnableTokenCache bool
	TokenCache       Cache
	Header           http.Header
	UserAgent        string
	Timeout          time.Duration
}

// RequestOption represents an option that can be modify an http.Request.
type RequestOption func(req *http.Request)

// WithHeader sets the provided header.
func WithHeader(header http.Header) RequestOption {
	return func(req *http.Request) {
		for k, vs := range header {
			for _, v := range vs {
				req.Header.Add(k, v)
			}
		}
	}
}

func WithAccessToken(token string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(hdrAuthorizationKey, token)
	}
}

func (c *Client) Request(ctx context.Context, method, apiPath string, body interface{}, v interface{}, options ...RequestOption) (*Response, error) {
	req, err := c.NewRequest(ctx, method, apiPath, body, options...)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, c.config.HttpClient, v)
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// which will be resolved to the BaseURL of the Client.
func (c *Client) NewRequest(ctx context.Context, method, apiPath string, body interface{}, opts ...RequestOption) (*http.Request, error) {
	var (
		buf     io.ReadWriter
		baseURL *url.URL
	)

	if c.config.BaseURL == nil {
		return nil, errors.New("BaseURL must be non-empty")
	}
	baseURL = c.config.BaseURL

	if body != nil {
		bs, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(bs)
	}

	newPath, err := baseURL.Parse(apiPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, newPath.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set(hdrContentTypeKey, contentTypeJSON)
	}
	if c.config.UserAgent != "" {
		req.Header.Set(hdrUserAgentKey, c.config.UserAgent)
	}
	if c.config.EnableTokenCache {
		token, _ := c.config.TokenCache.Get(ctx, "")
		opts = append(opts, WithAccessToken(token))
	}

	for _, opt := range opts {
		opt(req)
	}

	for k, vs := range c.config.Header {
		for _, v := range vs {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, httpClient *http.Client, v interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	if httpClient == nil {
		httpClient = c.config.HttpClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if er, ok := err.(*url.Error); ok {
			if url, err := url.Parse(er.URL); err == nil {
				er.URL = sanitizeURL(url).String()
				return nil, er
			}

			if er.Timeout() {
				return nil, fmt.Errorf("http: client timeout: %v", er.Error())
			}

			if e, ok := er.Err.(*net.OpError); ok && e.Op == "dial" {
				return nil, fmt.Errorf("http: client dial failed: %v", er.Error())
			}
		}
		return nil, err
	}

	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{Response: resp, RawBody: bs}

	var apiResp ApiResp
	err = json.Unmarshal(response.RawBody, &apiResp)
	if err != nil {
		return nil, err
	}

	if err := apiResp.CheckResult(); err != nil {
		return nil, err
	}

	return response, err
}

// sanitizeURL redacts the client parameter from the URL which may be
// exposed to the user.
func sanitizeURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	// u.User = nil
	// if u.Path == "" {
	// 	u.Path = "/"
	// }
	return u
}

type ClientOption func(conf *Config)

func WithEnableTokenCache(enableTokenCache bool) ClientOption {
	return func(config *Config) {
		config.EnableTokenCache = enableTokenCache
	}
}

func WithTokenCache(cache Cache) ClientOption {
	return func(config *Config) {
		config.TokenCache = cache
	}
}

func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(config *Config) {
		config.HttpClient = httpClient
	}
}

func WithHeaders(header http.Header) ClientOption {
	return func(config *Config) {
		config.Header = header
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(config *Config) {
		config.UserAgent = userAgent
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(config *Config) {
		config.Timeout = timeout
	}
}

func NewClient(appId int, secret string, opts ...ClientOption) *Client {
	adBaseURL, _ := url.Parse(baseURL)

	config := &Config{
		BaseURL:   adBaseURL,
		AppId:     appId,
		AppSecret: secret,
	}

	for _, opt := range opts {
		opt(config)
	}

	if config.HttpClient == nil {
		if config.Timeout == 0 {
			config.HttpClient = http.DefaultClient
		} else {
			config.HttpClient = &http.Client{Timeout: config.Timeout}
		}
	}

	c := &Client{config: config}

	c.common.client = c
	c.Auth = (*AuthService)(&c.common)
	c.Account = (*AccountService)(&c.common)
	c.Kol = (*KolService)(&c.common)
	c.Note = (*NoteService)(&c.common)
	c.Campaign = (*CampaignService)(&c.common)
	c.Unit = (*UnitService)(&c.common)
	c.Creativity = (*CreativityService)(&c.common)
	c.Report = (*ReportService)(&c.common)
	return c
}

type Response struct {
	*http.Response
	// Body is the response body.
	RawBody []byte `json:"-"`
}

var emptyPlaceholderRegx = regexp.MustCompile(`^\"-\"$`)

func (c *Client) JSONUnmarshalBody(r *Response, v interface{}) error {
	if !strings.Contains(r.Header.Get(hdrContentTypeKey), contentTypeJSON) {
		return fmt.Errorf("response content-type not json, response: %v", r)
	}
	str := strings.ReplaceAll(string(r.RawBody), "%", "")
	str = string(emptyPlaceholderRegx.ReplaceAll([]byte(str), []byte(`null`)))
	str = strings.ReplaceAll(str, `"-"`, `"0"`)
	if str == "" || str == "null" || str == "{}" || str == "[]" {
		return fmt.Errorf("response invaild or empty, response:%v", v)
	}
	err := json.Unmarshal([]byte(str), v)
	if err != nil {
		return err
	}
	return nil
}

// ApiResp 请求API响应通用字段
type ApiResp struct {
	Code      int    `json:"code"`                 // 返回码
	Msg       string `json:"msg"`                  // 返回信息
	Success   bool   `json:"success"`              // 接口是否成功
	RequestId string `json:"request_id,omitempty"` // 请求Id
}

func (a *ApiResp) IsSuccess() bool {
	return a.Success && (a.Code == 0 || a.Code == 200)
}

func (a *ApiResp) CheckResult() error {
	if a.IsSuccess() {
		return nil
	}
	return fmt.Errorf("code: %d, msg: %s", a.Code, a.Msg)
}

// PageResp 分页信息
type PageResp struct {
	PageIndex  int `json:"page_index"`  // 页码
	TotalCount int `json:"total_count"` // 总数量
}

// ListOptions specifies the optional parameters to various List methods that
// support offset pagination.
type ListOptions struct {
	AdvertiserId int64  `json:"advertiser_id"`       // 是	广告主Id
	PageNum      int    `json:"page_num,omitempty"`  // 否	页数，默认1
	PageSize     int    `json:"page_size,omitempty"` // 否	页大小，默认20,最大100
	StartDate    string `json:"start_date"`          // 是	开始时间，格式 yyyy-MM-dd	示例：2023-09-20
	EndDate      string `json:"end_date"`            // 是	结束时间，格式 yyyy-MM-dd	示例：2023-09-21
}
