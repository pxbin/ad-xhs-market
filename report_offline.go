package adxhsmarket

import (
	"context"
	"net/http"
)

// OfflineSearchWordRequest 表示创意层级离线数据请求参数
type OfflineSearchWordRequest struct {
	TimeUnit        string `json:"time_unit,omitempty"`        // 时间维度： "DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	MarketingTarget []int  `json:"marketing_target,omitempty"` // 营销诉求过滤条件： 3-商品推广_日常推广 4-产品种草 8-直播推广_日常推广 9-客资收集 10-抢占赛道 14-直播推广_直播预告 15-商品推广_店铺拉新 16-应用推广_应用唤起 17-外溢种草
	BiddingStrategy []int  `json:"bidding_strategy,omitempty"` // 出价方式：2：手动出价，101: 自动出价
	OptimizeTarget  []int  `json:"optimize_target,omitempty"`  // 推广目标过滤条件：0：点击量 1：互动量 3：表单提交量 4：商品成单量 5：私信咨询量 6：直播间观看量 11：商品访客量 12：落地页访问量 13：私信开口量 14：有效观看量 18：站外转化量 20：TI人群规模 21：行业商品成单 23：直播预热量 24：直播间成交 25：直播间支付ROI 35:APP打开（唤起） 36:APP进店（唤起） 37:APP互动（唤起） 38:APP成交-订单数（唤起） 43:APP打开按钮点击量
	Placement       []int  `json:"placement,omitempty"`        // 广告类型过滤条件：1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	PromotionTarget []int  `json:"promotion_target,omitempty"` // 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	Programmatic    []int  `json:"programmatic,omitempry"`     // 创意组合方式过滤条件 0：自定义创意 1：程序化创意
	BuildType       []int  `json:"build_type,omitempty"`       // 搭建方式过滤条件 0：标准搭建 1：省心智投
	SortColumn      string `json:"sort_column,omitempty"`      // 排序字段
	Sort            string `json:"sort,omitempty"`             // 升降序asc：升序desc：降序
	DataCaliber     int    `json:"data_caliber,omitempty"`     // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// OfflineSearchWordData 表示聚光数据报表创意层级离线数据响应
type OfflineSearchWordData struct {
	Page            PageRespDTO          `json:"page"`
	DataList        []OfflineCreativeDTO `json:"data_list"`
	AggregationData OfflineCreativeDTO   `json:"aggregation_data"`
}

// ListOfflineSearchWord 获取定向层级离线数据
func (s *ReportService) ListOfflineSearchWord(ctx context.Context, req *OfflineSearchWordRequest, options ...RequestOption) (*OfflineSearchWordData, error) {
	path := "/api/open/jg/data/report/offline/search/word"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[OfflineSearchWordData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OfflineAccountRequest 表示账户层级离线数据请求参数
type OfflineAccountRequest struct {
	TimeUnit        string   `json:"time_unit,omitempty"`        // 时间维度： "DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	MarketingTarget []int    `json:"marketing_target,omitempty"` // 营销诉求过滤条件： 3-商品推广_日常推广 4-产品种草 8-直播推广_日常推广 9-客资收集 10-抢占赛道 14-直播推广_直播预告 15-商品推广_店铺拉新 16-应用推广_应用唤起 17-外溢种草
	BiddingStrategy []int    `json:"bidding_strategy,omitempty"` // 出价方式：2：手动出价，101: 自动出价
	OptimizeTarget  []int    `json:"optimize_target,omitempty"`  // 推广目标过滤条件：0：点击量 1：互动量 3：表单提交量 4：商品成单量 5：私信咨询量 6：直播间观看量 11：商品访客量 12：落地页访问量 13：私信开口量 14：有效观看量 18：站外转化量 20：TI人群规模 21：行业商品成单 23：直播预热量 24：直播间成交 25：直播间支付ROI 35:APP打开（唤起） 36:APP进店（唤起） 37:APP互动（唤起） 38:APP成交-订单数（唤起） 43:APP打开按钮点击量
	Placement       []int    `json:"placement,omitempty"`        // 广告类型过滤条件：1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	PromotionTarget []int    `json:"promotion_target,omitempty"` // 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	BuildType       []int    `json:"build_type,omitempty"`       // 搭建方式过滤条件 0：标准搭建 1：省心智投
	SplitColumns    []string `json:"split_columns,omitempty"`    // 细分条件(相当于group by) 商品、落地页、直播间笔记只能三选一 	marketingTarget：营销诉求 buildType：搭建方式 placement：广告类型 optimizeTarget：推广目标 biddingStrategy：出价方式 promotionTarget：推广标的类型 jumpType：创意跳转类型 itemId：商品 pageId：落地页 liveRedId：直播间笔记 keyword：关键词 countryName: 国家（该字段数据于5月8号之后生效，之前日期不追溯暂无数据） province：省份 city：城市
	SortColumn      string   `json:"sort_column,omitempty"`      // 排序字段
	Sort            string   `json:"sort,omitempty"`             // 升降序asc：升序desc：降序
	DataCaliber     int      `json:"data_caliber,omitempty"`     // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// OfflineAccountData 表示聚光数据报表账户层级离线数据响应
type OfflineAccountData struct {
	DataList        []OfflineDataDTO `json:"data_list"`
	AggregationData OfflineDataDTO   `json:"aggregation_data"`
}

// ListOfflineAccount 获取定向层级离线数据
func (s *ReportService) ListOfflineAccount(ctx context.Context, req *OfflineAccountRequest, options ...RequestOption) (*OfflineAccountData, error) {
	path := "/api/open/jg/data/report/offline/account"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[OfflineAccountData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OfflineCampaignRequest 表示计划层级离线数据请求参数
type OfflineCampaignRequest struct {
	TimeUnit        string   `json:"time_unit,omitempty"`        // 时间维度： "DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	MarketingTarget []int    `json:"marketing_target,omitempty"` // 营销诉求过滤条件： 3-商品推广_日常推广 4-产品种草 8-直播推广_日常推广 9-客资收集 10-抢占赛道 14-直播推广_直播预告 15-商品推广_店铺拉新 16-应用推广_应用唤起 17-外溢种草
	BiddingStrategy []int    `json:"bidding_strategy,omitempty"` // 出价方式：2：手动出价，101: 自动出价
	OptimizeTarget  []int    `json:"optimize_target,omitempty"`  // 推广目标过滤条件：0：点击量 1：互动量 3：表单提交量 4：商品成单量 5：私信咨询量 6：直播间观看量 11：商品访客量 12：落地页访问量 13：私信开口量 14：有效观看量 18：站外转化量 20：TI人群规模 21：行业商品成单 23：直播预热量 24：直播间成交 25：直播间支付ROI 35:APP打开（唤起） 36:APP进店（唤起） 37:APP互动（唤起） 38:APP成交-订单数（唤起） 43:APP打开按钮点击量
	Placement       []int    `json:"placement,omitempty"`        // 广告类型过滤条件：1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	PromotionTarget []int    `json:"promotion_target,omitempty"` // 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	BuildType       []int    `json:"build_type,omitempty"`       // 搭建方式过滤条件 0：标准搭建 1：省心智投
	SplitColumns    []string `json:"split_columns,omitempty"`    // 细分条件(相当于group by) 商品、落地页、直播间笔记只能三选一 	marketingTarget：营销诉求 buildType：搭建方式 placement：广告类型 optimizeTarget：推广目标 biddingStrategy：出价方式 promotionTarget：推广标的类型 jumpType：创意跳转类型 itemId：商品 pageId：落地页 liveRedId：直播间笔记 keyword：关键词 countryName: 国家（该字段数据于5月8号之后生效，之前日期不追溯暂无数据） province：省份 city：城市
	SortColumn      string   `json:"sort_column,omitempty"`      // 排序字段
	Sort            string   `json:"sort,omitempty"`             // 升降序asc：升序desc：降序
	DataCaliber     int      `json:"data_caliber,omitempty"`     // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// OfflineCampaignData 表示聚光数据报表计划层级离线数据响应
type OfflineCampaignData struct {
	DataList        []OfflineDataDTO `json:"data_list"`
	AggregationData OfflineDataDTO   `json:"aggregation_data"`
}

// ListOfflineCampaign 获取定向层级离线数据
func (s *ReportService) ListOfflineCampaign(ctx context.Context, req *OfflineCampaignRequest, options ...RequestOption) (*OfflineCampaignData, error) {
	path := "/api/open/jg/data/report/offline/campaign"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[OfflineCampaignData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OfflineUnitRequest 表示单元层级离线数据请求参数
type OfflineUnitRequest struct {
	TimeUnit        string   `json:"time_unit,omitempty"`        // 时间维度： "DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	MarketingTarget []int    `json:"marketing_target,omitempty"` // 营销诉求过滤条件： 3-商品推广_日常推广 4-产品种草 8-直播推广_日常推广 9-客资收集 10-抢占赛道 14-直播推广_直播预告 15-商品推广_店铺拉新 16-应用推广_应用唤起 17-外溢种草
	BiddingStrategy []int    `json:"bidding_strategy,omitempty"` // 出价方式：2：手动出价，101: 自动出价
	OptimizeTarget  []int    `json:"optimize_target,omitempty"`  // 推广目标过滤条件：0：点击量 1：互动量 3：表单提交量 4：商品成单量 5：私信咨询量 6：直播间观看量 11：商品访客量 12：落地页访问量 13：私信开口量 14：有效观看量 18：站外转化量 20：TI人群规模 21：行业商品成单 23：直播预热量 24：直播间成交 25：直播间支付ROI 35:APP打开（唤起） 36:APP进店（唤起） 37:APP互动（唤起） 38:APP成交-订单数（唤起） 43:APP打开按钮点击量
	Placement       []int    `json:"placement,omitempty"`        // 广告类型过滤条件：1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	PromotionTarget []int    `json:"promotion_target,omitempty"` // 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	BuildType       []int    `json:"build_type,omitempty"`       // 搭建方式过滤条件 0：标准搭建 1：省心智投
	SplitColumns    []string `json:"split_columns,omitempty"`    // 细分条件(相当于group by) 商品、落地页、直播间笔记只能三选一 	marketingTarget：营销诉求 buildType：搭建方式 placement：广告类型 optimizeTarget：推广目标 biddingStrategy：出价方式 promotionTarget：推广标的类型 jumpType：创意跳转类型 itemId：商品 pageId：落地页 liveRedId：直播间笔记 keyword：关键词 countryName: 国家（该字段数据于5月8号之后生效，之前日期不追溯暂无数据） province：省份 city：城市
	SortColumn      string   `json:"sort_column,omitempty"`      // 排序字段
	Sort            string   `json:"sort,omitempty"`             // 升降序asc：升序desc：降序
	DataCaliber     int      `json:"data_caliber,omitempty"`     // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// OfflineUnitData 表示聚光单元层级离线报表数据
type OfflineUnitData struct {
	DataList        []OfflineDataDTO `json:"data_list"`
	AggregationData OfflineDataDTO   `json:"aggregation_data"`
}

// ListOfflineUnit 获取定向层级离线数据
func (s *ReportService) ListOfflineUnit(ctx context.Context, req *OfflineUnitRequest, options ...RequestOption) (*OfflineUnitData, error) {
	path := "/api/open/jg/data/report/offline/unit"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result, err := unmarshalApiResult[OfflineUnitData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OfflineCreativtyRequest 表示创意层级离线数据请求参数
type OfflineCreativtyRequest struct {
	TimeUnit        string `json:"time_unit,omitempty"`        // 时间维度： "DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	MarketingTarget []int  `json:"marketing_target,omitempty"` // 营销诉求过滤条件： 3-商品推广_日常推广 4-产品种草 8-直播推广_日常推广 9-客资收集 10-抢占赛道 14-直播推广_直播预告 15-商品推广_店铺拉新 16-应用推广_应用唤起 17-外溢种草
	BiddingStrategy []int  `json:"bidding_strategy,omitempty"` // 出价方式：2：手动出价，101: 自动出价
	OptimizeTarget  []int  `json:"optimize_target,omitempty"`  // 推广目标过滤条件：0：点击量 1：互动量 3：表单提交量 4：商品成单量 5：私信咨询量 6：直播间观看量 11：商品访客量 12：落地页访问量 13：私信开口量 14：有效观看量 18：站外转化量 20：TI人群规模 21：行业商品成单 23：直播预热量 24：直播间成交 25：直播间支付ROI 35:APP打开（唤起） 36:APP进店（唤起） 37:APP互动（唤起） 38:APP成交-订单数（唤起） 43:APP打开按钮点击量
	Placement       []int  `json:"placement,omitempty"`        // 广告类型过滤条件：1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	PromotionTarget []int  `json:"promotion_target,omitempty"` // 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	Programmatic    []int  `json:"programmatic,omitempry"`     // 创意组合方式过滤条件 0：自定义创意 1：程序化创意
	// // Deprecated 2024-09-02
	//BuildType       []int                  `json:"build_type,omitempty"`       // 搭建方式过滤条件 0：标准搭建 1：省心智投
	SplitColumns []string       `json:"split_columns,omitempty"` // 细分条件(相当于group by) 商品、落地页、直播间笔记只能三选一 	marketingTarget：营销诉求 buildType：搭建方式 placement：广告类型 optimizeTarget：推广目标 biddingStrategy：出价方式 promotionTarget：推广标的类型 jumpType：创意跳转类型 itemId：商品 pageId：落地页 liveRedId：直播间笔记 keyword：关键词 countryName: 国家（该字段数据于5月8号之后生效，之前日期不追溯暂无数据） province：省份 city：城市
	SortColumn   string         `json:"sort_column,omitempty"`   // 排序字段
	Sort         string         `json:"sort,omitempty"`          // 升降序asc：升序desc：降序
	DataCaliber  int            `json:"data_caliber,omitempty"`  // 否	数据指标归因时间类型 0-点击时间 1-转化时间
	Filters      []FilterClause `json:"filters,omitempty"`       // 过滤条件
	ListOptions
}

// OfflineCreativeData 表示聚光数据报表创意层级离线数据响应
type OfflineCreativeData struct {
	DataList        []OfflineCreativeDTO `json:"data_list"`
	AggregationData OfflineCreativeDTO   `json:"aggregation_data"`
}

// ListOfflineCreativty 获取定向层级离线数据
func (s *ReportService) ListOfflineCreativty(ctx context.Context, req *OfflineCreativtyRequest, options ...RequestOption) (*OfflineCreativeData, error) {
	path := "/api/open/jg/data/report/offline/creative"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[OfflineCreativeData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OfflinekeywordRequest 表示关键词层级离线数据请求参数
type OfflinekeywordRequest struct {
	Timekeyword     string   `json:"time_unit,omitempty"`        // 时间维度： "DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	MarketingTarget []int    `json:"marketing_target,omitempty"` // 营销诉求过滤条件： 3-商品推广_日常推广 4-产品种草 8-直播推广_日常推广 9-客资收集 10-抢占赛道 14-直播推广_直播预告 15-商品推广_店铺拉新 16-应用推广_应用唤起 17-外溢种草
	BiddingStrategy []int    `json:"bidding_strategy,omitempty"` // 出价方式：2：手动出价，101: 自动出价
	OptimizeTarget  []int    `json:"optimize_target,omitempty"`  // 推广目标过滤条件：0：点击量 1：互动量 3：表单提交量 4：商品成单量 5：私信咨询量 6：直播间观看量 11：商品访客量 12：落地页访问量 13：私信开口量 14：有效观看量 18：站外转化量 20：TI人群规模 21：行业商品成单 23：直播预热量 24：直播间成交 25：直播间支付ROI 35:APP打开（唤起） 36:APP进店（唤起） 37:APP互动（唤起） 38:APP成交-订单数（唤起） 43:APP打开按钮点击量
	Placement       []int    `json:"placement,omitempty"`        // 广告类型过滤条件：1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	PromotionTarget []int    `json:"promotion_target,omitempty"` // 投放标的类型：1: 笔记，2: 商品，7: 外链落地页，9: 落地页，18: 直播间
	BuildType       []int    `json:"build_type,omitempty"`       // 搭建方式过滤条件 0：标准搭建 1：省心智投
	SplitColumns    []string `json:"split_columns,omitempty"`    // 细分条件(相当于group by) 商品、落地页、直播间笔记只能三选一 	marketingTarget：营销诉求 buildType：搭建方式 placement：广告类型 optimizeTarget：推广目标 biddingStrategy：出价方式 promotionTarget：推广标的类型 jumpType：创意跳转类型 itemId：商品 pageId：落地页 liveRedId：直播间笔记 keyword：关键词 countryName: 国家（该字段数据于5月8号之后生效，之前日期不追溯暂无数据） province：省份 city：城市
	SortColumn      string   `json:"sort_column,omitempty"`      // 排序字段
	Sort            string   `json:"sort,omitempty"`             // 升降序asc：升序desc：降序
	DataCaliber     int      `json:"data_caliber,omitempty"`     // 否	数据指标归因时间类型 0-点击时间 1-转化时间

	ListOptions
}

// OfflinekeywordData 表示聚光关键词层级离线报表数据
type OfflinekeywordData struct {
	DataList        []OfflineKeywordDataDTO `json:"data_list"`
	TotalCount      int64                   `json:"total_count"`
	AggregationData OfflineKeywordDataDTO   `json:"aggregation_data"`
}

// ListOfflinekeyword 获取定向层级离线数据
func (s *ReportService) ListOfflinekeyword(ctx context.Context, req *OfflinekeywordRequest, options ...RequestOption) (*OfflinekeywordData, error) {
	path := "/api/open/jg/data/report/offline/keyword"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result, err := unmarshalApiResult[OfflinekeywordData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}
