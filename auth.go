package adxhsmarket

import (
	"context"
	"net/http"
)

type AuthService service

// AccessTokenRequest 表示开放平台授权Token请求
type AccessTokenRequest struct {
	AppId     int32  `json:"app_id"`
	AppSecret string `json:"secret"`
	AuthCode  string `json:"auth_code"`
}

// Advertiser 表示授权广告主账号，品牌开发者与服务商开发者时有值
type Advertiser struct {
	Id   int    `json:"advertiser_id"`   // 账号id
	Name string `json:"advertiser_name"` // 账号名称
}

// AccessTokenData 表示开放平台授权Token响应
type AccessTokenData struct {
	AccessToken           string       `json:"access_token"`             // 用于验证权限的 token
	AccessTokenExpiresIn  int          `json:"access_token_expires_in"`  // access_token 剩余有效时间，单位：秒
	RefreshToken          string       `json:"refresh_token"`            // 用于获取新的 access_token 和 refresh_token，并且刷新过期时间
	RefreshTokenExpiresIn int          `json:"refresh_token_expires_in"` // refresh_token 剩余有效时间，单位：秒
	UserId                string       `json:"user_id"`                  // 授权账号的user_id
	AppId                 int          `json:"app_id"`                   // 应用Id
	ApprovalRoleType      int          `json:"approval_role_type"`       // 授权账号类型，4：品牌，601：代理商
	RoleType              int          `json:"role_type"`                // 应用角色类型，1：品牌开发者，2：代理商开发者，3：服务商开发者
	PlatformType          int          `json:"platform_type"`            // 平台类型，1：聚光，2：蒲公英
	ApprovalAdvertisers   []Advertiser `json:"approval_advertisers"`     // 授权广告主账号，品牌开发者与服务商开发者时有值
	AdvertiserId          int          `json:"advertiser_id"`            // 授权账号账号id
	Scope                 string       `json:"scope"`                    // 授权接口范围
	CorporationName       string       `json:"corporation_name"`         // 企业名称
	CreateTime            int64        `json:"create_time"`              // 授权记录创建时间
	UpdateTime            int64        `json:"update_time"`              // 授权记录更新时间
}

type AccessTokenResponse struct {
	ApiResp
	Data AccessTokenData `json:"data"`
}

// AccessToken 获取Oauth2.0 token
func (s *AuthService) AccessToken(ctx context.Context, authCode string, options ...RequestOption) (*AccessTokenResponse, error) {
	path := "/api/open/oauth2/access_token"
	body := map[string]interface{}{"app_id": s.client.config.AppId, "secret": s.client.config.AppSecret, "auth_code": authCode}

	response, err := s.client.Request(ctx, http.MethodPost, path, body, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &AccessTokenResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RefreshToken Token有效期&续期
func (s *AuthService) RefreshToken(ctx context.Context, token string, options ...RequestOption) (*AccessTokenResponse, error) {
	path := "/api/open/oauth2/refresh_token"
	body := map[string]interface{}{"app_id": s.client.config.AppId, "secret": s.client.config.AppSecret, "refresh_token": token}

	response, err := s.client.Request(ctx, http.MethodPost, path, body, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &AccessTokenResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
