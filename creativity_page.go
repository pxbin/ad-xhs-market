package adxhsmarket

import (
	"context"
	"net/http"
)

type CreativityPageCreateRequest struct {
	AdvertiserId       int64              `json:"advertiser_id"`                   // 是	广告主Id
	UnitId             int64              `json:"unit_id"`                         // 是	单元Id
	CreativityName     string             `json:"creativity_name,omitempty"`       // 是	创意名称
	PageCreativityInfo PageCreativityInfo `json:"page_creativity_infos,omitempty"` // 否 前链H5
}

// CreativityPageCreate 创建落地页创意
func (s *CreativityService) CreativityPageCreate(ctx context.Context, req *CreativityNoteCreateRequest, options ...RequestOption) (*CreateCreativityResponse, error) {
	path := "/api/open/jg/creativity/page/create"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &CreateCreativityResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
