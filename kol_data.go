package adxhsmarket

import (
	"context"
	"net/http"
)

type KolService service

type KolInfo struct {
	KolNickName     string `json:"kol_nick_name"`
	RedId           string `json:"red_id"`
	KolFanNum       int64  `json:"kol_fan_num"`
	KolId           string `json:"kol_id"`
	KolOperateState int    `json:"kol_operate_state"`
	KolCreditLevel  int    `json:"kol_credit_level"`
	McnName         string `json:"mcn_name"`
	VideoPrice      int64  `json:"video_price"`
	Price           int64  `json:"price"`
}

type KolDetailData struct {
	DateKey string    `json:"date_key"`
	Items   []KolInfo `json:"datas"`
}

type ListKolDetailsResponse struct {
	ApiResp
	Data NotePostData `json:"data"`
}

func (s *KolService) ListKolDetails(ctx context.Context, userId string, kolIds []string, options ...RequestOption) (*ListKolDetailsResponse, error) {
	path := "/api/open/pgy/kol/data/detail"
	body := map[string]interface{}{"user_id": userId, "kol_ids": kolIds}

	response, err := s.client.Request(ctx, http.MethodPost, path, body, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &ListKolDetailsResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
