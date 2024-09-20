package adxhsmarket

import (
	"context"
	"net/http"
)

// CampaignService 表示推广计划服务
type CampaignService service

type (
	// TimePeriod 自定义时间段
	TimePeriod struct {
		Mon  string `json:"mon"`  //	星期一	默认24个1：111111111111111111111111示例：101111111111111111111111每个小时用0和1表示，0表示不投，1表示投放，示例中表示1点不投，其他时间投以下星期同理
		Tues string `json:"tues"` //	星期二	默认：111111111111111111111111
		Wed  string `json:"wed"`  //	星期三	默认：111111111111111111111111
		Thur string `json:"thur"` //	星期四	默认：111111111111111111111111
		Fri  string `json:"fri"`  //	星期五	默认：111111111111111111111111
		Sat  string `json:"sat"`  //	星期六	默认：111111111111111111111111
		Sun  string `json:"sun"`  //	星期日	默认：111111111111111111111111
	}

	// CampaignIdDTO 计划id
	CampaignIdDTO struct {
		CampaignId int64 `json:"campaign_id"`
	}

	PageReqDTO struct {
		PageIndex int `json:"page_index,omitempty"` // 否	页码	默认1
		PageSize  int `json:"page_size,omitempty"`  // 否	每页查询量级	默认20，
	}
)

type CampaignCreateRequest struct {
	AdvertiserId           int64       `json:"advertiser_id"`                      // 是	广告主Id
	MarketingTarget        int         `json:"marketing_target"`                   // 是	营销目标（4：产品种草，9：客资收集，16：应用推广）
	CampaignName           string      `json:"campaign_name"`                      // 是	计划名称（长度不超过50个字符）
	Placement              int         `json:"placement"`                          // 是	广告类型（1：信息流，2：搜索推广））
	PromotionTarget        int         `json:"promotion_target"`                   // 是	推广标的类型，（1：笔记）
	Enable                 int         `json:"enable,omitempty"`                   // 否	计划创建后默认开启状态，1-开启，0-不开启，不传默认开启
	TimeType               int         `json:"time_type"`                          // 是	推广时时间类型,，0:长期投放，1:自定义设置开始结束时间
	StartTime              string      `json:"start_time,omitempty"`               // 否	推广开始时间，格式 yyyy-MM-dd	示例：2023-09-20，长期投放的开始时间可不填，自定义的时间设置需在今天及以后
	ExpireTime             string      `json:"expire_time,omitempty"`              // 否	推广结束时间，格式 yyyy-MM-dd	示例：2023-09-21，长期投放的结束时间可不填，自定义的设置时间需在今天及以后
	TimePeriodType         int         `json:"time_period_type"`                   // 是	推广时段类型, 0: 全时段，1:自定义时间段
	TimePeriod             *TimePeriod `json:"time_period,omitempty"`              // 否	高级设置-自定义时间段（全时段可不传）
	BiddingStrategy        int         `json:"bidding_strategy"`                   // 是	出价方式 2: 手动出价 3: 自动出价-成本自动控制（BCB） 4:自动出价（MCB） 7: 自动出价-成本手动控制（OCPX）	自动出价-成本手动控制时设置（即bidding_strategy=7）唤端场景下传7
	LimitDayBudget         int         `json:"limit_day_budget"`                   // 是	预算类型，0：不限预算，1：指定预算
	CampaignDayBudget      int         `json:"campaign_day_budget,omitempty"`      // 否	计划日预算，单位分，范围 [10000~99999900)，>=10000, < 99999900	指定预算时设置
	OptimizeTarget         int         `json:"optimize_target"`                    // 是 推广目标
	ConstraintType         int         `json:"constraint_type,omitempty"`          // 否 成本控制类型
	SmartSwitch            int         `json:"smart_switch,omitempty"`             // 否	节假日预算上浮, 0:关闭，1:开启	不限预算不支持节假日上浮(不传或传0都可）
	PacingMode             int         `json:"pacing_mode,omitempty"`              // 否	投放速率, 1:匀速，2:加速	不限预算可不填，默认是加速，自动出价-成本自动控制，固定匀速
	FeedFlag               int         `json:"feed_flag,omitempty"`                // 否	搜索追投	是否开启搜索追投：（搜索推广-普通投放下，互动成本控制和手动出价下支持搜索追投功能）0：关闭1：开启
	BuildType              int         `json:"build_type,omitempty"`               // 否	搭建类型	0：普通搭建 （标准搭建）1：智能搭建 （省心智投）默认为0（普通搭建），唤端下支持智能搭建
	EventAssetId           int64       `json:"event_asset_id,omitempty"`           // 否	资产id。	默认0见/api/open/jg/data/event/asset/info 接口返回。这里对应返回值的event_asset_id.唤端下需要关注
	AssetEvent             int64       `json:"asset_event,omitempty"`              // 否	资产事件类型，这里和推广目标一一对应。401：APP打开（唤起）402：APP进店（唤起）403：APP互动（唤起）404：APP成交-订单数（唤起）	默认0。见/api/open/jg/data/event/asset/info 接口返回。这里对应返回值的event_type.唤端下需要关注
	AssetEventId           int64       `json:"asset_event_id,omitempty"`           // 否	资产事件id	默认0见/api/open/jg/data/event/asset/info 接口返回。这里对应返回值的event_id.唤端下需要关注
	PageCategory           int         `json:"page_category,omitempty"`            // 否	落地页类型	默认01：聚光落地页2：自研落地页3：原生落地页
	SearchFlag             int         `json:"search_flag,omitempty"`              // 否	搜索快投	0：开启后关闭，1：开启
	TargetExtensionSwatich int         `json:"target_extension_swatich,omitempty"` // 否	搜索快投定向拓展	默认0
	SearchBidRatio         float32     `json:"search_bid_ratio,omitempty"`         // 否	搜索快投-出价系数	默认1.0
	DeeplinkId             int64       `json:"deeplink_id,omitempty"`              // 否	deeplink链接id	唤端场景下必填
	UniversalLinkId        int64       `json:"universal_link_id,omitempty"`        // 否	ulk的链接id	唤端场景下需要，非必填
	DetectURLLink          string      `json:"detect_url_link,omitempty"`          // 否	监测链接	唤端场景下需要关注。optimize_target如果是35、36、37、38，则必填
}

type CampaignIdData struct {
	CampaignId int64 `json:"campaign_id"` // 计划id
}

type CreateCampaignResponse struct {
	ApiResp
	Data CampaignIdData `json:"data"`
}

func (s *CampaignService) Create(ctx context.Context, req *CampaignCreateRequest, options ...RequestOption) (*CreateCampaignResponse, error) {
	path := "/api/open/jg/campaign/create"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &CreateCampaignResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

type CampaignUpdateRequest struct {
	AdvertiserId          int64       `json:"advertiser_id"`                     // 是	广告主Id
	MarketingTarget       int         `json:"marketing_target"`                  // 是	营销目标（4：产品种草，9：客资收集，16：应用推广）
	CampaignId            int64       `json:"campaign_id"`                       // 是	计划Id
	CampaignName          string      `json:"campaign_name"`                     // 是	计划名称（长度不超过50个字符）
	LimitDayBudget        int         `json:"limit_day_budget"`                  // 是	预算类型，0：不限预算，1：指定预算
	CampaignDayBudget     int         `json:"campaign_day_budget,omitempty"`     // 否	计划日预算，单位分，范围 [10000~99999900)，>=10000, < 99999900	指定预算时设置
	SmartSwitch           int         `json:"smart_switch,omitempty"`            // 否	节假日预算上浮, 0:关闭，1:开启	不限预算不支持节假日上浮(不传或传0都可）
	TimeType              int         `json:"time_type"`                         // 是	推广时时间类型,，0:长期投放，1:自定义设置开始结束时间
	StartTime             string      `json:"start_time,omitempty"`              // 否	推广开始时间，格式 yyyy-MM-dd	示例：2023-09-20，长期投放的开始时间可不填，自定义的时间设置需在今天及以后
	ExpireTime            string      `json:"expire_time,omitempty"`             // 否	推广结束时间，格式 yyyy-MM-dd	示例：2023-09-21，长期投放的结束时间可不填，自定义的设置时间需在今天及以后
	TimePeriodType        int         `json:"time_period_type"`                  // 是	推广时段类型, 0: 全时段，1:自定义时间段
	TimePeriod            *TimePeriod `json:"time_period,omitempty"`             // 否	高级设置-自定义时间段（全时段可不传）
	PacingMode            int         `json:"pacing_mode,omitempty"`             // 否	投放速率, 1:匀速，2:加速	不限预算可不填，默认是加速，自动出价-成本自动控制，固定匀速
	FeedFlag              int         `json:"feed_flag,omitempty"`               // 否	搜索追投	是否开启搜索追投：（搜索推广-普通投放下，互动成本控制和手动出价下支持搜索追投功能）0：关闭1：开启
	BiddingStrategy       int         `json:"bidding_strategy"`                  // 是	出价方式 2: 手动出价 3: 自动出价-成本自动控制（BCB） 4:自动出价（MCB） 7: 自动出价-成本手动控制（OCPX）	自动出价-成本手动控制时设置（即bidding_strategy=7）唤端场景下传7
	SearchFlag            int         `json:"search_flag,omitempty"`             // 否	搜索快投	0：开启后关闭，1：开启
	TargetExtensionSwtich int         `json:"target_extension_switch,omitempty"` // 否	搜索快投定向拓展	默认0
	SearchBidRatio        float32     `json:"search_bid_ratio,omitempty"`        // 否	搜索快投-出价系数	默认1.0
	DeeplinkId            int64       `json:"deeplink_id,omitempty"`             // 否	deeplink链接id	唤端场景下必填
	UniversalLinkId       int64       `json:"universal_link_id,omitempty"`       // 否	ulk的链接id	唤端场景下需要，非必填
	DetectURLLink         string      `json:"detect_url_link,omitempty"`         // 否	监测链接	唤端场景下需要关注。optimize_target如果是35、36、37、38，则必填
}

type UpdateCampaignResponse struct {
	ApiResp
	Data CampaignIdData `json:"data"`
}

func (s *CampaignService) Update(ctx context.Context, req *CampaignUpdateRequest, options ...RequestOption) (*UpdateCampaignResponse, error) {
	path := "/api/open/jg/campaign/update"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &UpdateCampaignResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

type UpdateCampaignStatusRequest struct {
	AdvertiserId int64   `json:"advertiser_id"` // 是	广告主Id	必填
	CampaignIds  []int64 `json:"campaign_ids"`  // 是	计划Id列表	至少传一个限制单次变更计划数量，最多传20
	ActionType   int     `json:"action_type"`   //是	操作类型	1：开启2：暂停3：删除
}

type UpdateCampaignStatusData struct {
	CampaignIds []int64 `json:"campaign_ids"`
}

type UpdateCampaignStatusResponse struct {
	ApiResp
	Data UpdateCampaignStatusData `json:"data"`
}

func (s *CampaignService) UpdateStatus(ctx context.Context, req *UpdateCampaignStatusRequest, options ...RequestOption) (*UpdateCampaignStatusResponse, error) {
	path := "/api/open/jg/campaign/status/update"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &UpdateCampaignStatusResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

type ListCampaignRequest struct {
	AdvertiserId int64      `json:"advertiser_id"`         // 是	广告主Id	必填
	CampaignIds  []int64    `json:"campaign_ids"`          // 是	计划Id列表	至少传一个限制单次变更计划数量，最多传20
	StartTime    string     `json:"start_time,omitempty"`  // 否	开始时间	示例：2023-09-20，和expire_time配套填写，创建时间查询范围-开始时间
	ExpireTime   string     `json:"expire_time,omitempty"` // 否	结束时间	示例：2023-09-21，start_time配套填写，创建时间查询范围-结束时间
	Status       int        `json:"status,omitempty"`      // 否	创意状态	1-有效 2-暂停 3-已删除 4-计划预算不足，5-现金余额不足，7-账户日预算不足 8-处于暂停阶段
	Page         PageReqDTO `json:"page"`
}

type CampaignData struct {
	CampaignId          int64   `json:"campaign_id"`           // 计划Id
	CampaignName        int     `json:"campaign_name"`         // 计划名称
	CampaignFilterState int     `json:"campaign_filter_state"` // 计划状态	1-有效，2-暂停，3-已删除，4-计划预算不足，5-现金余额不足，6-所有未删除状态，7-账户日预算不足
	CampaignCreateTime  string  `json:"campaign_create_time"`  // 计划创建时间
	CampaignEnable      int     `json:"campaign_enable"`       // 计划是否可用
	MarketingTarget     int     `json:"marketing_target"`      // 营销目标	3-商品销量4-产品种草8-直播推广9-客资收集10-抢占关键词13-种草直达14-直播预热15-店铺拉新 16-应用推广
	Placement           int     `json:"placement"`             // 广告类型	1-信息流2-搜索推广4-全站智投7-视频内流
	OptimizeTarget      int     `json:"optimize_target"`       // 优化目标
	PromotionTarget     int     `json:"promotion_target"`      // 投放标的
	BiddingStrategy     int     `json:"bidding_strategy"`      // 出价策略
	ConstraintType      int     `json:"constraint_type"`       // 成本控制类型
	ConstraintValue     int     `json:"constraint_value"`      // 成本值
	LimitDayBudget      int     `json:"limit_day_budget"`      // 预算类型	0-不限预算，1-指定预算
	CampaignDayBudget   int     `json:"campaign_day_budget"`   // 计划日预算
	BudgetState         int     `json:"budget_state"`          // 推广计划日预算是否充足，	0-不足，1-充足
	SmartSwitch         int     `json:"smart_switch"`          // 智能开关
	Platform            int     `json:"platform"`              // 创建来源
	PacingMode          int     `json:"pacing_mode"`           // 投放速率	1-匀速2-加速
	StartTime           string  `json:"start_time"`            // 推广开始时间
	ExpireTime          string  `json:"expire_time"`           // 推广结束时间
	TimePeriod          string  `json:"time_period"`           // 推广时间的bitmap
	TimePeriodType      int     `json:"time_period_type"`      // 推广时间段类型	0-全时段1-自定义时间段
	FeedFlag            int     `json:"feed_flag"`             // 是否开启搜索追投,	0-关1-开
	BuildType           int     `json:"build_type"`            // 构建类型
	CreativityState     int     `json:"creativity_state"`      // 总预算达成时间	创意聚合状态，ark电商推广场景使用
	EventAssetId        int64   `json:"event_asset_id"`        // 事件资产
	AssetEvent          int64   `json:"asset_event"`           // 资产事件
	AssetEventId        int64   `json:"asset_event_id"`        // 资产事件id
	PageCategory        int     `json:"page_category"`         // 落地页类型
	SearchFlag          int     `json:"search_flag"`           // 搜索快投开关
	SearchBidRatio      float64 `json:"search_bid_ratio"`      // 定向拓展
	DeeplinkId          int64   `json:"deeplink_id"`           // 唤端场景下的deeplink id
	UniversalLinkId     int64   `json:"universal_link_id"`     // 唤端场景下的universalLink
	DetectURLLink       string  `json:"detect_url_link"`       // 唤端场景下的监测链接
}

type ListCampaignData struct {
	Page      PageReqDTO     `json:"page"`
	Campaigns []CampaignData `json:"base_campaign_dtos"`
}

type ListCampaignResponse struct {
	ApiResp
	Data ListCampaignData `json:"data"`
}

func (s *CampaignService) List(ctx context.Context, req *ListCampaignRequest, options ...RequestOption) (*ListCampaignResponse, error) {
	path := "/api/open/jg/campaign/list"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &ListCampaignResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
