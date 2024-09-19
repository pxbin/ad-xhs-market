package adxhsmarket

import (
	"context"
	"net/http"
)

// CreativityService 表示推广创意的服务
type CreativityService service

type (

	// CreativityIdData 创意Id
	CreativityIdData struct {
		Id int64 `json:"creativity_id"`
	}

	// CreativityIdsData 创意Id集合
	CreativityIdsData struct {
		Ids []int64 `json:"creativity_id"`
	}

	// QualInfo 资质信息
	QualInfo struct {
		ApplyId           string `json:"apply_id"`            // 申请id，对应行业资质	是
		ProductQualIdList []int  `json:"product_qualId_list"` // 产品资质id	否
		BrandQualIdList   []int  `json:"brand_qualId_list"`   // 品牌资质id	否
	}

	// H5Info 创意图片信息
	H5Info struct {
		PhotoURL       string   `json:"photo_url"`                 // 是	图片链接url	格式仅支持PNG、JPG、JPEG图片大小:最大值4M
		Content        string   `json:"content"`                   // 是	文案	<= 25字
		ClickUrls      []string `json:"click_urls,omitempty"`      // 否	点击链接
		ExpoUrls       []string `json:"expo_urls,omitempty"`       // 否	曝光链接
		MonitorCompany string   `json:"monitor_company,omitempty"` // 否	监测公司
		MonitorParams  string   `json:"monitor_params,omitempty"`  // 否	监测参数配置
	}

	// PageCreativityInfo 前链h5信息
	PageCreativityInfo struct {
		PageId   string   `json:"page_id"`      // 落地页id
		H5Infos  []H5Info `json:"h5_Info_dtos"` // 创意图片信息
		QualInfo QualInfo `json:"qual_info"`    // 资质信息
	}

	// Photo 创意图片
	Photo struct {
		PhotoUrl string `json:"photo_url"` //
	}

	// Title 标题
	Title struct {
		Title string `json:"title"` // 标题
	}

	// H5MaterialInfo 程序化创意信息
	H5MaterialInfo struct {
		Photos         []Photo  `json:"photos"`          // 是	创意图片
		Titles         []Title  `json:"titles"`          // 是	标题
		QualInfo       QualInfo `json:"qual_info"`       // 是	资质信息
		ClickUrls      []string `json:"click_urls"`      // 否	点击链接
		ExpoUrls       []string `json:"expo_urls"`       // 否	曝光链接
		MonitorCompany string   `json:"monitor_company"` // 否	监测公司
		MonitorParams  string   `json:"monitor_params"`  // 否	监测参数
	}
)

type CreativityUpdateRequest struct {
	AdvertiserId            int64          `json:"advertiser_id"`                       // 是	广告主id
	CreativityId            int64          `json:"creativity_id"`                       // 是	创意id
	CreativityName          string         `json:"creativity_name,omitempty"`           // 否	创意名称
	ClickUrls               []string       `json:"click_urls,omitempty"`                // 否	点击链接
	ExpoUrls                []string       `json:"expo_urls,omitempty"`                 // 否	曝光链接
	MaskPerfer              bool           `json:"mask_perfer,omitempty"`               // 否	是否开启封面优选
	TitleMaskPerfer         bool           `json:"title_mask_perfer,omitempty"`         // 否	是否开启标题优选
	JumpURL                 string         `json:"jump_url,omitempty"`                  // 否	跳转链接，包括自有链接和落地页链接
	BarContent              string         `json:"bar_content,omitempty"`               // 否	文案内容
	ItemId                  string         `json:"item_id,omitempty"`                   // 否	商品id
	H5Infos                 H5Info         `json:"h5_infos,omitempty"`                  // 否	前链h5信息
	ConversionComponentType []int          `json:"conversion_component_type,omitempty"` // 否	转化组件类型	0-营销组件1-评论区组件
	Comment                 string         `json:"comment,omitempty"`                   // 否	评论区文案
	H5MaterialInfo          H5MaterialInfo `json:"h5_material_info,omitempty"`          // 否	程序化创意素材
	PoiId                   string         `json:"poi_id,omitempty"`                    // 否	poiId
	PoiJumpType             string         `json:"poi_jump_type,omitempty"`             // 否	poi组件跳转类型
	MonitorCompany          string         `json:"monitor_company,omitempty"`           // 否
	MonitorParams           string         `json:"monitor_params,omitempty"`            // 否	监测参数配置
	AdBizItemId             string         `json:"ad_biz_item_id,omitempty"`            // 否	广告侧绑定商品id
	AppCompIcon             string         `json:"app_comp_icon,omitempty"`             // 否	唤端下，商品主图
	FallBackJumpUrl         string         `json:"fall_back_jump_url,omitempty"`        // 否	唤端下，兜底链接
}

// Update 编辑创意
func (s *CreativityService) Update(ctx context.Context, req *CreativityUpdateRequest, options ...RequestOption) (*interface{}, error) {
	path := "/api/open/jg/creativity/update"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[interface{}](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UpdateCreativityStatusRequest struct {
	AdvertiserId  int64   `json:"advertiser_id"`  // 是	广告主Id	必填
	CreativityIds []int64 `json:"creativity_ids"` // 是	创意id列表 最多传20
	ActionType    int     `json:"action_type"`    //是	操作类型	1：开启2：暂停3：删除
}

type UpdateCreativityStatusData struct {
	CampaignIds []int64 `json:"campaign_ids"`
}

func (s *CreativityService) UpdateStatus(ctx context.Context, req *UpdateCampaignStatusRequest, options ...RequestOption) (*CreativityIdsData, error) {
	path := "/api/open/jg/creativity/status/update"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[CreativityIdsData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ListCreativityRequest struct {
	AdvertiserId  int64      `json:"advertiser_id"`         // 是	广告主Id
	CampaignId    int64      `json:"campaign_id,omitempty"` // 否	计划Id
	UnitId        int64      `json:"unit_id,omitempty"`     // 否	单元Id
	CreativityIds []int64    `json:"creativity_ids"`        // 否 创意 Id集合
	Status        int        `json:"status,omitempty"`      // 否	创意状态	1 已删除2 所有未删除状态3 暂停4 已被单元暂停5 已被计划暂停8 有效9 商品状态异常10 单元未开始11 单元已结束12 单元处于暂停时段
	StartTime     string     `json:"start_time,omitempty"`  // 否	开始时间	示例：2023-09-20，和expire_time配套填写，创建时间查询范围-开始时间
	ExpireTime    string     `json:"expire_time,omitempty"` // 否	结束时间	示例：2023-09-21，start_time配套填写，创建时间查询范围-结束时间
	Page          PageReqDTO `json:"page"`
}

type CreativityData struct {
	AdvertiserId             int64       `json:"advertiser_id"`              // 广告主id
	CampaignId               int64       `json:"campaign_id"`                // 计划id
	UnitId                   int64       `json:"unit_id"`                    // 单元id
	CreativityId             int64       `json:"creativity_id"`              // 创意id
	CreativityName           string      `json:"creativity_name"`            // 创意名称
	CreativityEnable         int         `json:"creativity_enable"`          // 创意开启状态
	CreativityFilterState    int         `json:"creativity_filter_state"`    // 创意状态	1 已删除2 所有未删除状态3 暂停4 已被单元暂停5 已被计划暂停8 有效9 商品状态异常10 单元未开始11 单元已结束12 单元处于暂停时段 13-计划预算不足 14-现金余额不足 16-账户日预算不足
	CreativityCreateTime     string      `json:"creativity_create_time"`     // 创意创建时间
	MaterialType             int         `json:"material_type"`              // 笔记类型
	ConversionType           int         `json:"conversion_type"`            // 组件类型，	0: 无组件，1：商品组件2: 落地页组件，3: 私信组件，4: 直播组件5: poi门店组件6: 外链商品7: 直播间8: 搜索组件9: 小程序组件10：留资组件11：唤端组件（新功能，后续支持）
	NoteId                   string      `json:"note_id"`                    // 笔记id
	NoteType                 int         `json:"note_type"`                  // 笔记类型
	CustomMask               int         `json:"custom_mask"`                // 自选封面，仅固定首图为封面传2，笔记全部图片均可作为封面传1，（默认0）	智能创意相关字段，支持控制原笔记的图片是否可参与创意优选
	CustomTitle              int         `json:"custom_title"`               // 是否自提标题 是传1，否传2，（默认0）	智能创意相关字段，支持自定义更多的标题来参与创意优选
	TitleFills               []string    `json:"title_fills"`                // 自提标题，和笔记原标题共计上限10条。	是否自提标题为“2”时，传了无效
	MaskGen                  int         `json:"mask_gen"`                   // 是否开启自动优化封面 是传1，否传2，（默认0）	智能创意相关字段，开启后，系统将通过对人群特征和海量搜索词的学习自动优化封面，使展现的封面更贴近用户兴趣意图，从而提升投放效果。
	TitleGen                 int         `json:"title_gen"`                  // 是否开启自动优化标题 是传1，否传2，（默认0）	智能创意相关字段，开启后，系统将通过对人群特征和海量搜索词的学习自动优化标题，使展现的标题更贴近用户兴趣意图，从而提升投放效果。
	MaskPrefer               bool        `json:"mask_prefer"`                // 是否开启封面优选
	TitleMaskPrefer          bool        `json:"title_mask_prefer"`          // 是否开启标题优选
	AuditStatus              int         `json:"audit_status"`               // 审核状态
	AuditComment             string      `json:"audit_comment"`              // 审核备注（驳回原因）
	PageId                   string      `json:"page_id"`                    // 落地页id
	ClickUrls                []string    `json:"click_urls"`                 // 点击链接
	ExpoUrls                 []string    `json:"expo_urls"`                  // 曝光链接
	JumpUrl                  string      `json:"jump_url"`                   // 落地页/外链url
	BarContent               string      `json:"bar_content"`                // 按钮文案内容	落地页组件类型：立即参与、立即购买、立即领取、立即预约；私信组件类型：立即咨询、立即参与、立即领取、立即预约；
	Image                    string      `json:"image"`                      // 创意图片
	ItemInvalidReason        int         `json:"item_invalid_reason"`        // 商品状态异常原因
	ConversionComponentTypes []int64     `json:"conversion_component_types"` //组件位置，0：默认位置，1：置顶评论
	Comment                  string      `json:"comment"`                    // 置顶评论文案	当conversion_component_types包含置顶评论时，评论文案
	Programmatic             int         `json:"programmatic"`               // 是否是程序化创意	0-非程序化创意1-程序化创意
	CreativityExtraInfo      string      `json:"creativity_extra_info"`      // 创意extrainfo
	IntoShopParam            string      `json:"into_shop_param"`            // 店铺小程序转化组件填写参数
	BootScreenInfo           interface{} `json:"boot_screen_info"`           // 开屏创意
	PoiId                    string      `json:"poi_id"`                     // poiId
	PoiJumpType              string      `json:"poi_jump_type"`              // poi组件跳转链接
	MonitorCompany           string      `json:"monitor_company"`            // 监测公司
	MonitorParams            string      `json:"monitor_params"`             // 监测参数配置
	ItemId                   string      `json:"item_id"`                    // 商品id
	Title                    string      `json:"title"`                      // 直播标题	直播推广-标的为直播间时 推广标的为直播时选
	GoodsSellingPoint        string      `json:"goods_selling_point"`        // 商品卖点
	DataPostUrl              string      `json:"data_post_url"`              // 数据回传url
	KosMsgType               int         `json:"kos_msg_type"`               // kos私信承接方	0-私信tok1-私信tob
	QualInfo                 QualInfo    `json:"qual_info"`                  // 资质信息
	MiniProgramPath          string      `json:"mini_program_path"`          // 小程序组件path
}

type ListCreativityData struct {
	Page         PageRespDTO      `json:"page"`
	Creativities []CreativityData `json:"creativity_dtos"`
}

func (s *CreativityService) List(ctx context.Context, req *ListCampaignRequest, options ...RequestOption) (*ListCreativityData, error) {
	path := "/api/open/jg/creativity/search"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[ListCreativityData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}
