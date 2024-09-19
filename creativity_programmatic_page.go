package adxhsmarket

import (
	"context"
	"net/http"
)

type ProgrammaticCreateRequest struct {
	AdvertiserId   int64          `json:"advertiser_id"`              // 是	广告主Id
	UnitId         int64          `json:"unit_id"`                    // 是	单元Id
	CreativityName string         `json:"creativity_name,omitempty"`  // 是	创意名称
	H5MaterialInfo H5MaterialInfo `json:"h5_material_info,omitempty"` // 否 前链H5
}

// ProgrammaticCreate 创建程序化创意
func (s *CreativityService) ProgrammaticCreate(ctx context.Context, req *ProgrammaticCreateRequest, options ...RequestOption) (*CreativityIdData, error) {
	path := "/api/open/jg/creativity/programmatic/page/create"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[CreativityIdData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}
