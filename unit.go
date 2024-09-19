package adxhsmarket

import (
	"context"
	"net/http"
)

// UnitService 表示推广单元服务
type UnitService service

type (
	Unit struct {
		AdvertiserId        int64            `json:"advertiser_id"`                    // 是	广告主Id
		CampaignId          int64            `json:"campaign_id"`                      // 是	计划Id
		UnitName            string           `json:"unit_name"`                        //是	单元名称
		EventBid            int              `json:"event_bid,omitempty"`              // 否	出价/目标成本，单位分，自动控制不需要传	校验见下面出价校验部分
		NoteIds             []string         `json:"note_ids"`                         // 是	标的-笔记id(帮助定向推荐，创意里的笔记会单独填写)非种草人群规模、深度种草人群规模下必填
		PromotionTarget     int              `json:"promotion_target,omitempty"`       // 否	推广标的（笔记1）
		TargetType          int              `json:"target_type"`                      // 是	定向类型，1-通投,2-智能定向, 3-高级定向
		TargetConfig        TargetConfig     `json:"target_config"`                    //是	定向配置
		KeywordTargetPeriod int              `json:"keyword_target_period,omitempty"`  // 否	关键词时间周期，单位天，枚举包括 3，7，15，30
		KeywordTargetAction []int            `json:"keyword_target_action,omitempty"`  // 否	关键词行为类型，1: 搜索，2: 互动，3: 阅读
		BusinessTreeName    string           `json:"business_tree_name,omitempty"`     // 否	推广业务信息示例：生活服务>婚纱摄影;美妆个护;母婴>母婴食品>奶粉具体参看/api/open/jg/keyword/industry/taxonomy接口返回词。	不能跨一级行业
		SpuNoteInfo         []SpuNoteConfig  `json:"spu_note_info,omitempty"`          // 否	spu&笔记标的信息推广目标是种草人群规模、深度种草人群规模下必传。spu只能一个，至少绑定5篇笔记，最多50(api限制)只能选择合作笔记与我的笔记，笔记违规不能选，spu审核不通过也可以选，选择的笔记一定是绑定在该spu上。
		KeywordWithBid      []KeywordWithBid `json:"keyword_with_bid,omitempty"`       // 否	关键词，通过(关键词词包)/api/open/jg/keyword/word/bag/list (智能推词、行业推词、以词推词)/api/open/jg/keyword/common/recommend 接口获取搜索推广-标准投放下必传
		SubstitutedUserId   string           `json:"substituted_user_id,omitempty"`    // 否	代投账号b的userId，代投笔记与其他笔记互斥,，需要校验是否有代投账号权限
		KeywordGenType      int              `json:"keyword_gen_type,omitempty"`       // 否	单元选词方式： -1:无意义默认值 0:手动选词 1:智能拓词 2:手动+智能 关键词定向且在白名单中才支持
		PageId              string           `json:"page_id,omitempty"`                // 否	落地页Id，聚光落地页下必填	客资收集下支持/api/open/jg/landing_page/list_landing_page接口object_id字段
		LandingPageURL      string           `json:"landing_page_url,omitempty"`       // 否	落地页Url，自研落地页下必填	客资收集下支持
		UnitExternalPageURL string           `json:"unit_external_page_url,omitempty"` // 否	外链Url，标的是外链落地页时必填	客资收集下支持
		UnitLandingPageDesc []string         `json:"unit_landing_page_desc,omitempty"` // 否	落地页表单描述	客资收集下支持
		TargetTemplateId    int64            `json:"target_template_id,omitempty"`     // 否	定向包id
	}

	// UnitIdDTO 计划id
	UnitIdDTO struct {
		UnitIdDTO int64 `json:"unit_id"`
	}

	SpuNoteConfig struct {
		SpuId   string   `json:"spu_id"`   // 是	spuid
		NoteIds []string `json:"note_ids"` // 是	绑定的笔记id
	}

	KeywordWithBid struct {
		Keyword         string `json:"keyword"`           // 关键词
		Bid             int    `json:"bid"`               // 关键词出价，手动出价时在0.3-10元之间，自动出价时为0。
		KeywordSource   int    `json:"keyword_source"`    // 关键词来源，手动加词可以不传
		PhraseMatchType int    `json:"phrase_match_type"` // 匹配方式0:精确匹配, 1:短语匹配
		FeedBid         int    `json:"feed_bid"`          // 搜索追投出价,开启搜索追投时必填，0.3-10元
	}

	TargetConfig struct {
		TargetGender           string                 `json:"target_gender,omitempty"`             // 否	性别，不限: all, 男: 0, 女: 1 在(种草、客资或者唤端营销诉求) 并且（信息流、视频内流、全站智投）下必须要传 其他场景不限制
		TargetAge              string                 `json:"target_age,omitempty"`                // 否	年龄，不限：all, 细分年龄段：18-22, 23-27, 28-32, 33-100，细分年龄后多个年龄段用#号分隔，如18-22#23-27 非搜索下必传。 搜索下，需要配置白名单，白名单里的必传，白名单外的不能传。
		TargetCity             string                 `json:"target_city"`                         // 是	地域定向-城市，不限传 all	需要传二级城市，不可传省份，字典见下面定向的通用接口
		TargetDevice           string                 `json:"target_device,omitempty"`             // 否	设备，不限: all, 苹果: ios, 安卓: android 在(种草、客资或者唤端营销诉求) 并且（信息流、搜索或者视频内流）必传； 搜索下 并且在搜索白名单里，必传； 其他情况下可以不传
		IndustryInterestTarget IndustryInterestTarget `json:"industry_interest_target,omitempty"`  // 否	兴趣定向	值来自于定向信息接口的返回的industry_interest_target
		CrowdTarget            CrowdTarget            `json:"crowd_target"`                        // 否	dmp人群包定向	值来自于定向信息接口的返回的crowd_target
		InterestKeywords       []string               `json:"interest_keywords,omitempty"`         // 否	关键词兴趣
		Keywords               []string               `json:"keywords,omitempty"`                  // 否	关键词行为
		IntelligentExpansion   int                    `json:"intelligent_expansion,omitempty"`     // 否	智能扩量，0: 否，1: 是
		SearchTargetCityIntent string                 `json:"search_target_city_intent,omitempty"` // 否	搜索地域意图定向功能  0-关闭，1-开启
	}

	IndustryInterestTarget struct {
		ContentInterests  []CodeNamePair `json:"content_interests,omitempty"`  // 行业阅读兴趣 值来自于定向信息接口的返回的industry_interest_target#content_interests只有一级
		ShoppingInterests []CodeNamePair `json:"shopping_Interests,omitempty"` // 行业购物兴趣（需要填写至二级）	值来自于定向信息接口的返回的industry_interest_target#shopping_Interests需要一级和二级结构
	}

	CodeNamePair struct {
		Code     string         `json:"code"`     // code
		Name     string         `json:"name"`     // 名称
		Children []CodeNamePair `json:"children"` // 子节点
	}

	CrowdTarget struct {
		CrowdPkg []CrowdPackage `json:"crowd_pkg"` //
	}

	CrowdPackage struct {
		Value string `json:"value"` //	人群包Id
		Name  string `json:"name"`  //	人群包名称
		Type  string `json:"type"`  // 人群包类型：common/outside:平台精选-场景人群，timeliness: 节促人群，空: 人群包
	}

	ItemNoteConfig struct {
		ItemId  string   `json:"item_id"`  // 是	spuid
		NoteIds []string `json:"note_ids"` // 是	绑定的笔记id
	}

	RecommendTarget struct {
		HighPotential         string `json:"high_potential"`          // 高潜词包
		InterestHighPotential string `json:"interest_high_potential"` // 关键词兴趣-高潜词包
	}

	DandelionCrowd struct {
		NormalDandelionCrowdList     []NormalDandelionCrowd     `json:"normal_dandelion_crowd_list"`
		CustomizedDandelionCrowdList []CustomizedDandelionCrowd `json:"customized_dandelion_crowd_List"`
	}

	NormalDandelionCrowd struct {
		ActionType  string `json:"action_type"`   // 行为类型:imp,read
		TimePeriod  int    `json:"time_period"`   // 时间周期，30，90
		BrandUserId string `json:"brand_user_id"` // 品牌用户Id
	}

	CustomizedDandelionCrowd struct {
		ActionType string   `json:"action_type"`  // 行为类型:imp,read
		TimePeriod int      `json:"time_period"`  // 时间周期，30，90
		CrowdId    int64    `json:"crowd_id"`     // 蒲公英定制人群包Id
		CrowdName  string   `json:"crowd_name"`   // 蒲公英定制人群包名称
		NoteIdList []string `json:"note_id_list"` // 笔记ids
		Channels   []string `json:"channels"`     // 流量渠道：ad/nature
	}
)

type UnitCreateRequest struct {
	AdvertiserId        int64            `json:"advertiser_id"`                    // 是	广告主Id
	CampaignId          int64            `json:"campaign_id"`                      // 是	计划Id
	UnitName            string           `json:"unit_name"`                        //是	单元名称
	EventBid            int              `json:"event_bid,omitempty"`              // 否	出价/目标成本，单位分，自动控制不需要传	校验见下面出价校验部分
	NoteIds             []string         `json:"note_ids"`                         // 是	标的-笔记id(帮助定向推荐，创意里的笔记会单独填写)非种草人群规模、深度种草人群规模下必填
	PromotionTarget     int              `json:"promotion_target,omitempty"`       // 否	推广标的（笔记1）
	TargetType          int              `json:"target_type"`                      // 是	定向类型，1-通投,2-智能定向, 3-高级定向
	TargetConfig        TargetConfig     `json:"target_config"`                    //是	定向配置
	KeywordTargetPeriod int              `json:"keyword_target_period,omitempty"`  // 否	关键词时间周期，单位天，枚举包括 3，7，15，30
	KeywordTargetAction []int            `json:"keyword_target_action,omitempty"`  // 否	关键词行为类型，1: 搜索，2: 互动，3: 阅读
	BusinessTreeName    string           `json:"business_tree_name,omitempty"`     // 否	推广业务信息示例：生活服务>婚纱摄影;美妆个护;母婴>母婴食品>奶粉具体参看/api/open/jg/keyword/industry/taxonomy接口返回词。	不能跨一级行业
	SpuNoteInfo         []SpuNoteConfig  `json:"spu_note_info,omitempty"`          // 否	spu&笔记标的信息推广目标是种草人群规模、深度种草人群规模下必传。spu只能一个，至少绑定5篇笔记，最多50(api限制)只能选择合作笔记与我的笔记，笔记违规不能选，spu审核不通过也可以选，选择的笔记一定是绑定在该spu上。
	KeywordWithBid      []KeywordWithBid `json:"keyword_with_bid,omitempty"`       // 否	关键词，通过(关键词词包)/api/open/jg/keyword/word/bag/list (智能推词、行业推词、以词推词)/api/open/jg/keyword/common/recommend 接口获取搜索推广-标准投放下必传
	SubstitutedUserId   string           `json:"substituted_user_id,omitempty"`    // 否	代投账号b的userId，代投笔记与其他笔记互斥,，需要校验是否有代投账号权限
	KeywordGenType      int              `json:"keyword_gen_type,omitempty"`       // 否	单元选词方式： -1:无意义默认值 0:手动选词 1:智能拓词 2:手动+智能 关键词定向且在白名单中才支持
	PageId              string           `json:"page_id,omitempty"`                // 否	落地页Id，聚光落地页下必填	客资收集下支持/api/open/jg/landing_page/list_landing_page接口object_id字段
	LandingPageURL      string           `json:"landing_page_url,omitempty"`       // 否	落地页Url，自研落地页下必填	客资收集下支持
	UnitExternalPageURL string           `json:"unit_external_page_url,omitempty"` // 否	外链Url，标的是外链落地页时必填	客资收集下支持
	UnitLandingPageDesc []string         `json:"unit_landing_page_desc,omitempty"` // 否	落地页表单描述	客资收集下支持
	TargetTemplateId    int64            `json:"target_template_id,omitempty"`     // 否	定向包id
}

func (s *UnitService) Create(ctx context.Context, req *UnitCreateRequest, options ...RequestOption) (*UnitIdDTO, error) {
	path := "/api/open/jg/unit/create"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[UnitIdDTO](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UnitUpdateRequest struct {
	AdvertiserId        int64            `json:"advertiser_id"`                    // 是	广告主Id
	CampaignId          int64            `json:"campaign_id"`                      // 是	计划Id
	UnitId              int64            `json:"unit_id"`                          // 是	单元名称
	UnitName            string           `json:"unit_name"`                        // 是	单元名称
	EventBid            int              `json:"event_bid,omitempty"`              // 否	出价/目标成本，单位分，自动控制不需要传	校验见下面出价校验部分
	NoteIds             []string         `json:"note_ids"`                         // 是	标的-笔记id(帮助定向推荐，创意里的笔记会单独填写)非种草人群规模、深度种草人群规模下必填
	PromotionTarget     int              `json:"promotion_target,omitempty"`       // 否	推广标的（笔记1）
	TargetType          int              `json:"target_type"`                      // 是	定向类型，1-通投,2-智能定向, 3-高级定向
	TargetConfig        TargetConfig     `json:"target_config"`                    //是	定向配置
	KeywordTargetPeriod int              `json:"keyword_target_period,omitempty"`  // 否	关键词时间周期，单位天，枚举包括 3，7，15，30
	KeywordTargetAction []int            `json:"keyword_target_action,omitempty"`  // 否	关键词行为类型，1: 搜索，2: 互动，3: 阅读
	BusinessTreeName    string           `json:"business_tree_name,omitempty"`     // 否	推广业务信息示例：生活服务>婚纱摄影;美妆个护;母婴>母婴食品>奶粉具体参看/api/open/jg/keyword/industry/taxonomy接口返回词。	不能跨一级行业
	SpuNoteInfo         []SpuNoteConfig  `json:"spu_note_info,omitempty"`          // 否	spu&笔记标的信息推广目标是种草人群规模、深度种草人群规模下必传。spu只能一个，至少绑定5篇笔记，最多50(api限制)只能选择合作笔记与我的笔记，笔记违规不能选，spu审核不通过也可以选，选择的笔记一定是绑定在该spu上。
	KeywordWithBid      []KeywordWithBid `json:"keyword_with_bid,omitempty"`       // 否	关键词，通过(关键词词包)/api/open/jg/keyword/word/bag/list (智能推词、行业推词、以词推词)/api/open/jg/keyword/common/recommend 接口获取搜索推广-标准投放下必传
	SubstitutedUserId   string           `json:"substituted_user_id,omitempty"`    // 否	代投账号b的userId，代投笔记与其他笔记互斥,，需要校验是否有代投账号权限
	KeywordGenType      int              `json:"keyword_gen_type,omitempty"`       // 否	单元选词方式： -1:无意义默认值 0:手动选词 1:智能拓词 2:手动+智能 关键词定向且在白名单中才支持
	PageId              string           `json:"page_id,omitempty"`                // 否	落地页Id，聚光落地页下必填	客资收集下支持/api/open/jg/landing_page/list_landing_page接口object_id字段
	LandingPageURL      string           `json:"landing_page_url,omitempty"`       // 否	落地页Url，自研落地页下必填	客资收集下支持
	UnitExternalPageURL string           `json:"unit_external_page_url,omitempty"` // 否	外链Url，标的是外链落地页时必填	客资收集下支持
	UnitLandingPageDesc []string         `json:"unit_landing_page_desc,omitempty"` // 否	落地页表单描述	客资收集下支持
	TargetTemplateId    int64            `json:"target_template_id,omitempty"`     // 否	定向包id
}

func (s *UnitService) Update(ctx context.Context, req *UnitUpdateRequest, options ...RequestOption) (*UnitIdDTO, error) {
	path := "/api/open/jg/unit/update"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[UnitIdDTO](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UpdateUnitStatusRequest struct {
	AdvertiserId int64   `json:"advertiser_id"` // 是	广告主Id	必填
	UnitIds      []int64 `json:"unit_ids"`      // 是	单元Id列表	至少传一个限制单次变更计划数量，最多传20
	Status       int     `json:"status"`        // 是	状态类型	1：开启2：暂停3：删除
}

type UpdateUnitStatusData struct {
	UnitIds []int64 `json:"unit_ids"`
}

func (s *UnitService) UpdateStatus(ctx context.Context, req *UpdateUnitStatusRequest, options ...RequestOption) (*UpdateUnitStatusData, error) {
	path := "/api/open/jg/unit/status/update"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[UpdateUnitStatusData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ListUnitRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`         // 是	广告主Id	必填
	CampaignId   int64   `json:"campaign_id,omitempty"` // 否	计划Id列表	至少传一个限制单次变更计划数量，最多传20
	UnitIds      []int64 `json:"unit_ids,omitempty"`    // 否	单元Id列表	至少传一个限制单次变更计划数量，最多传20
	Status       int     `json:"status,omitempty"`      // 否	创意状态	1-有效 2-暂停 3-已删除 4-计划预算不足，5-现金余额不足，7-账户日预算不足 8-处于暂停阶段
	UnitName     string  `json:"unit_name,omitempty"`   // 否	单元名称
	StartDate    string  `json:"start_date,omitempty"`  // 否	开始时间	示例：2023-09-20，和expire_time配套填写，创建时间查询范围-开始时间
	EndData      string  `json:"end_date,omitempty"`    // 否	结束时间	示例：2023-09-21，start_time配套填写，创建时间查询范围-结束时间
	Page         int     `json:"page,omitemty"`
	PageSize     int     `json:"page_size,omitemty"`
}

type UnitData struct {
	Id                  int64            `json:"id"`                     // 单元id
	CampaignId          int64            `json:"campaign_id"`            // 计划id
	Name                string           `json:"name"`                   // 单元名称
	Enable              int              `json:"enable"`                 // 投放状态：0：暂停1：投放中
	EventBid            int              `json:"event_bid"`              // 出价,单位(分)
	TargetType          int              `json:"target_type"`            // 定向类型1-通投,2-智能定向,3-高级定向
	ItemIds             []string         `json:"item_ids"`               // 商品Id
	NoteIds             []string         `json:"note_ids"`               // 笔记Id
	LiveUserId          string           `json:"live_user_id"`           // 直播用户Id
	PageId              string           `json:"page_id"`                // 落地页Id
	LandingPageURL      string           `json:"landing_page_url"`       // 落地页Url
	UnitExternalPageURL string           `json:"unit_external_page_url"` // 外链Url
	LandingPageType     int              `json:"landing_page_type"`      // 落地页链接类型:1-表单,2-外跳链接，0-默认值，无实际意义
	TargetPosition      int              `json:"target_position"`        // 抢占资源1-首位,3-第三位，0-不限位置
	TargetGoal          int              `json:"target_goal"`            // 抢占目标1-点击抢占市场份额，0-默认值，无实际意义
	WordTagName         string           `json:"word_tag_name"`          // 词包名称
	ProportionGoal      float64          `json:"proportion_goal"`        // 占比目标
	BusinessTreeName    string           `json:"business_tree_name"`     // 推广业务信息示例：生活服务>婚纱摄影;美妆个护;母婴>母婴食品>奶粉
	UnitLandingPageDesc []string         `json:"unit_landing_page_desc"` // 落地页表单描述
	KeywordTargetPeriod int              `json:"keyword_target_period"`  // 关键词定向行为周期，单位天，枚举包括 3，7，15，30
	KeywordTargetAction []int            `json:"keyword_target_action"`  // 关键词定向行为1:搜索,2:互动,3:阅读
	SubstitutedUserId   string           `json:"substituted_user_id"`    // 代投账号b的userId
	CreateTime          string           `json:"create_time"`            // 单元创建时间
	UpdateTime          string           `json:"update_time"`            // 单元修改时间
	ItemNoteInfo        []ItemNoteConfig `json:"item_note_info"`         // 单元标的信息
	SpuNoteInfo         []SpuNoteConfig  `json:"spu_note_info"`          // spu&笔记标的信息
	TargetConfig        TargetConfig     `json:"target_config"`          // 定向信息
	KeywordGenType      int              `json:"keyword_gen_type"`       //单元选词方式： -1:无意义默认值 0:手动选词 1:智能拓词 2:手动+智能
	KeywordWithBid      []KeywordWithBid `json:"keyword_with_bid"`       //
}

type ListUnitData struct {
	TotalCount int        `json:"total_count"` // 总数
	UnitInfos  []UnitData `json:"unit_infos"`  // 单元信息
}

func (s *UnitService) List(ctx context.Context, req *ListUnitRequest, options ...RequestOption) (*ListUnitData, error) {
	path := "/api/open/jg/campaign/list"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[ListUnitData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}
