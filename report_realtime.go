package adxhsmarket

import (
	"context"
	"net/http"
)

// TargetDTO 定向数据
type TargetDTO struct {
	Data             DataReportDTO   `json:"data"`
	BaseCampaignData BaseCampaignDTO `json:"base_campaign_dto"`
	BaseUnitData     BaseUnitDTO     `json:"base_unit_dto"`
	BaseTargetData   BaseTargetDTO   `json:"base_target_dto	"`
}

// RealtimeTargetResponse 表示聚光数据报表定向层级实时数据响应
type RealtimeTargetResponse struct {
	BaseResp
	Page       PageRespDTO   `json:"page"`        // 分页信息
	TotalData  DataReportDTO `json:"total_data"`  // 汇总数据
	TargetDTOs []TargetDTO   `json:"target_dtos"` // 定向数据
}

// RealtimeTargetRequest 表示定向层级实时数据请求参数
type RealtimeTargetRequest struct {
	SortColumn          string `json:"sort_column,omitempty"`           // 排序字段
	Sort                string `json:"sort,omitempty"`                  // 升降序asc：升序desc：降序
	Name                string `json:"name,omitempty"`                  // 搜索定向名称
	MarketingTargetList []int  `json:"marketing_target_list,omitempty"` // 营销诉求筛选
	ListOptions
}

// ListRealtimeTarget 获取创意层级实时数据
func (s *ReportService) ListRealtimeTarget(ctx context.Context, req *RealtimeTargetRequest, options ...RequestOption) (*RealtimeTargetResponse, error) {
	path := "/api/open/jg/data/report/realtime/target"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &RealtimeTargetResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RealtimeAccountRequest 表示账户层级实时数据请求参数
type RealtimeAccountRequest struct {
	NeedHourlyData bool `json:"need_hourly_data,omitempty"`

	ListOptions
}

// RealtimeAccountResponse 表示聚光数据报表账户层级实时数据响应
type RealtimeAccountResponse struct {
	BaseResp
	Data       DataReportDTO   `json:"data"`
	HourlyData []DataReportDTO `json:"hourly_data"` // 小时数据
}

// ListRealtimeAccount 获取账户层级实时数据
func (s *ReportService) ListRealtimeAccount(ctx context.Context, req *RealtimeAccountRequest, options ...RequestOption) (*RealtimeAccountResponse, error) {
	path := "/api/open/jg/data/report/realtime/account"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &RealtimeAccountResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RealtimeCampaignRequest 表示计划层级实时数据请求参数
type RealtimeCampaignRequest struct {
	SortColumn              string `json:"sort_column,omitempty"`                // 排序字段
	Sort                    string `json:"sort,omitempty"`                       // 升降序asc：升序desc：降序
	MarketingTargetList     []int  `json:"marketing_target_list,omitempty"`      // 营销诉求筛选
	CampaignFilterState     int    `json:"campaign_filter_state,omitempty"`      // 否	计划状态过滤，8：有效，3：暂停，9：商品状态异常，4：已被单元暂停，10：单元未开始，11：单元已结束，12：单元处于暂停时段，5：已被计划暂停，13：计划预算不足，16：账户日预算不足，14：现金余额不足，1：已删除，2：所有未删除状态，
	CampaignCreateBeginTime string `json:"campaign_create_begin_time,omitempty"` // 否	计划创建时间范围的开始
	CampaignCreateEndTime   string `json:"campaign_create_end_time,omitempty"`   // 否	计划创建时间范围的结束
	PlacementList           []int  `json:"placement_list,omitempty"`             // 否	广告类型：1：信息流广告，2：搜索广告，7：视频流广告
	LimitDayBudgetList      []int  `json:"limit_day_budget_list,omitempty"`      // 否	预算类型：0：不限预算，1：指定预算
	OptimizeTargetList      []int  `json:"optimize_target_list,omitempty"`       // 否	推广目标：0：点击量，1：互动量，16：种草值，11：商品访客量，12：落地页访问量，3：表单提交量，4：商品成单量，5：私信咨询量，6：观看量，13：私信开口量，14：有效观看量，17：ROI，站外转化量，20：TI人群规模 35:APP打开（唤起）36:APP进店（唤起）37:APP互动（唤起）38:APP成交-订单数（唤起）43:APP打开按钮点击量
	BuildTypeList           []int  `json:"build_type_list,omitempty"`            // 否	搭建方式：0：标准投放，1：省心智投
	BiddingStrategyList     []int  `json:"bidding_strategy_list,omitempty"`      // 否	出价方式：2：手动出价，101: 自动出价
	ConstraintTypeList      []int  `json:"constraint_type_list,omitempty"`       // 否	成本控制方式：-1: 无，101: 自动控制，0: 点击成本控制，1: 互动成本控制，3: 表单提交成本控制，5: 私信咨询成本控制，11: 访客成本控制，13: 私信开口成本控制，14: 有效观播成本控制，17: ROI控制，23: 预热成本控制，50: 私信留资成本控制
	PromotionTargetList     []int  `json:"promotion_target_list,omitempty"`      // 否	投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	CombineAuditStatus      int    `json:"combine_audit_status,omitempty"`       // 否	创意审核状态，1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密）
	MigrationStatusList     []int  `json:"migration_status_list,omitempty"`      // 否	计划迁移状态，0：非迁移计划，2：迁移计划	专业号平台的计划迁移至聚光
	Name                    string `json:"name,omitempty"`                       // 否	搜索计划名称
	Id                      int    `json:"id,omitempty"`                         // 否	搜索计划id
	DataCaliber             int    `json:"data_caliber,omitempty"`               // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// CampaignDTO  计划数据
type CampaignDTO struct {
	Data             DataReportDTO   `json:"data"`
	BaseCampaignData BaseCampaignDTO `json:"base_campaign_dto"` // 计划属性信息
}

// RealtimeCampaignResponse 表示聚光数据报表计划层级实时数据响应
type RealtimeCampaignResponse struct {
	BaseResp
	Page         PageRespDTO   `json:"page"`
	TotalData    DataReportDTO `json:"total_data"`    // 汇总数据
	CampaignDTOs []CampaignDTO `json:"campaign_dtos"` // 计划数据list
}

// ListRealtimeCampaign 获取定向层级实时数据
func (s *ReportService) ListRealtimeCampaign(ctx context.Context, req *RealtimeCampaignRequest, options ...RequestOption) (*RealtimeCampaignResponse, error) {
	path := "/api/open/jg/data/report/realtime/campaign"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &RealtimeCampaignResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RealtimeUnitRequest 表示单元层级实时数据请求参数
type RealtimeUnitRequest struct {
	SortColumn          string `json:"sort_column,omitempty"`            // 排序字段
	Sort                string `json:"sort,omitempty"`                   // 升降序asc：升序desc：降序
	MarketingTargetList []int  `json:"marketing_target_list,omitempty"`  // 营销诉求筛选
	UnitFilterState     int    `json:"unit_filter_state,omitempty"`      // 否	单元状态过滤，8：有效，3：暂停，9：商品状态异常，4：已被单元暂停，10：单元未开始，11：单元已结束，12：单元处于暂停时段，5：已被单元暂停，13：单元预算不足，16：账户日预算不足，14：现金余额不足，1：已删除，2：所有未删除状态，
	UnitCreateBeginTime string `json:"unit_create_begin_time,omitempty"` // 否	单元创建时间范围的开始
	UnitCreateEndTime   string `json:"unit_create_end_time,omitempty"`   // 否	单元创建时间范围的结束
	PlacementList       []int  `json:"placement_list,omitempty"`         // 否	广告类型：1：信息流广告，2：搜索广告，7：视频流广告
	BiddingStrategyList []int  `json:"bidding_strategy_list,omitempty"`  // 否	出价方式：2：手动出价，101: 自动出价
	PromotionTargetList []int  `json:"promotion_target_list,omitempty"`  // 否	投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	CombineAuditStatus  int    `json:"combine_audit_status,omitempty"`   // 否	创意审核状态，1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密）
	Name                string `json:"name,omitempty"`                   // 否	搜索单元名称
	Id                  int    `json:"id,omitempty"`                     // 否	搜索单元id
	DataCaliber         int    `json:"data_caliber,omitempty"`           // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// UnitDTO 单元数据
type UnitDTO struct {
	Data             DataReportDTO   `json:"data"`
	BaseUnitData     BaseUnitDTO     `json:"base_unit_dto"`     // 单元属性信息
	BaseCampaignData BaseCampaignDTO `json:"base_compaign_dto"` // 计划属性信息
}

// RealtimeUnitResponse 表示聚光数据报表单元层级实时数据响应
type RealtimeUnitResponse struct {
	BaseResp
	Page      PageRespDTO   `json:"page"`
	TotalData DataReportDTO `json:"total_data"` // 汇总数据
	UnitDTOs  []UnitDTO     `json:"unit_dtos"`  // 单元数据list
}

// ListRealtimeUnit 获取定向层级实时数据
func (s *ReportService) ListRealtimeUnit(ctx context.Context, req *RealtimeUnitRequest, options ...RequestOption) (*RealtimeUnitResponse, error) {
	path := "/api/open/jg/data/report/realtime/unit"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &RealtimeUnitResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RealtimeCreativityRequest 表示创意层级实时数据请求参数
type RealtimeCreativityRequest struct {
	SortColumn                string `json:"sort_column,omitempty"`                  // 排序字段
	Sort                      string `json:"sort,omitempty"`                         // 升降序asc：升序desc：降序
	PlacementList             []int  `json:"placement_list,omitempty"`               // 否	广告类型：1：信息流广告，2：搜索广告，7：视频流广告
	CreativityFilterState     int    `json:"creativity_filter_state,omitempty"`      // 否	创意状态过滤，8：有效，3：暂停，9：商品状态异常，4：已被单元暂停，10：单元未开始，11：单元已结束，12：单元处于暂停时段，5：已被计划暂停，13：计划预算不足，16：账户日预算不足，14：现金余额不足，1：已删除，2：所有未删除状态，
	CreativityCreateBeginTime string `json:"creativity_create_begin_time,omitempty"` // 否	创意创建时间范围的开始
	CreativityCreateEndTime   string `json:"creativity_create_end_time,omitempty"`   // 否	创意创建时间范围的结束
	ConversionType            int    `json:"conversion_type,omitempty"`              // 否	创意类型：30：商品，20：落地页，4：直播间笔记，7：直播间，0：笔记（无组件），1：笔记（商品组件），2：笔记（落地页组件）
	ProgrammaticList          []int  `json:"programmatic_list,omitempty"`            // 否	创意组合方式，0：自定义创意，1：程序化创意
	CreativityAuditState      int    `json:"creativity_audit_state,omitempty"`       // 否	创意审核状态，1：审核拒绝，2：审核中，3：审核通过，4：审核通过（私密）
	Name                      string `json:"name,omitempty"`                         // 否	搜索创意名称
	Id                        int64  `json:"id,omitempty"`                           // 否	搜索创意id
	DataCaliber               int    `json:"data_caliber,omitempty"`                 // 否	数据指标归因时间类型 0-点击时间 1-转化时间
	ListOptions
}

// CreativityDTO 创意数据
type CreativityDTO struct {
	BaseCreativityData BaseCreativityDTO `json:"base_creativity_dto"`
	BaseCampaignData   BaseCampaignDTO   `json:"base_campaign_dto"`
	Data               DataReportDTO     `json:"data"`
	BaseUnitData       BaseUnitDTO       `json:"base_unit_dto"`
}

// RealtimeCreativityResponse 表示聚光数据报表创意层级实时数据响应
type RealtimeCreativityResponse struct {
	BaseResp
	Page           PageRespDTO     `json:"page"`
	TotalData      DataReportDTO   `json:"total_data"` // 汇总数据
	CreativityDTOs []CreativityDTO `json:"creativity_dtos"`
}

// ListRealtimeCreativity 获取定向层级实时数据
func (s *ReportService) ListRealtimeCreativity(ctx context.Context, req *RealtimeCreativityRequest, options ...RequestOption) (*RealtimeCreativityResponse, error) {
	path := "/api/open/jg/data/report/realtime/creativity"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &RealtimeCreativityResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RealtimeKeywordRequest 表示关键词层级实时数据请求参数
type RealtimeKeywordRequest struct {
	SortColumn         string `json:"sort_column,omitempty"`          // 排序字段
	Sort               string `json:"sort,omitempty"`                 // 升降序asc：升序desc：降序
	KeywordFilterState int    `json:"keyword_filter_state,omitempty"` // 否	关键词状态过滤，8：有效，3：暂停，9：商品状态异常，4：已被关键词暂停，10：关键词未开始，11：关键词已结束，12：关键词处于暂停时段，5：已被关键词暂停，13：关键词预算不足，16：账户日预算不足，14：现金余额不足，1：已删除，2：所有未删除状态，
	UseBidStrategy     int    `json:"use_bid_strategy,omitempty"`     // 否	出价策略：0：未使用出价策略1：已使用出价策略
	KeywordName        string `json:"keyword_name,omitempty"`         // 否	搜索关键词名词
	CampaignName       string `json:"campaign_name	,omitempty"`       // 否	搜索计划名称
	UnitName           string `json:"unit_name,omitempty"`            // 否		搜索单元名称
	DataCaliber        int    `json:"data_caliber,omitempty"`         // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// KeywordDTO 关键词数据
type KeywordDTO struct {
	Data             DataReportDTO   `json:"data"`
	BaseCampaignData BaseCampaignDTO `json:"base_campaign_dto"`
	BaseUnitData     BaseUnitDTO     `json:"base_unit_dto"`
	BaseKeywordData  BaseKeywordDTO  `json:"base_keyword_dTO"` // 关键词属性信息
	SubKeywordDTOs   []KeywordDTO    `json:"sub_keyword_dtos"` // 子关键词信息
}

// RealtimeKeywordResponse 表示聚光数据报表关键词层级实时数据响应
type RealtimeKeywordResponse struct {
	BaseResp
	Page        PageRespDTO   `json:"page"`
	TotalData   DataReportDTO `json:"total_data"`   // 汇总数据
	KeywordDTOs []KeywordDTO  `json:"keyword_dtos"` // 关键词数据list
}

// ListRealtimeKeyword 获取定向层级实时数据
func (s *ReportService) ListRealtimeKeyword(ctx context.Context, req *RealtimeKeywordRequest, options ...RequestOption) (*RealtimeKeywordResponse, error) {
	path := "/api/open/jg/data/report/realtime/keyword"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &RealtimeKeywordResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
