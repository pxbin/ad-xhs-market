package adxhsmarket

import (
	"context"
	"net/http"
)

// NoteService 表示蒲公英头后数据服务
type NoteService service

type (

	// AudiencePersonaData 阅读受众分析
	AudiencePersonaData struct {
		FanPercent       float64 `json:"fan_percent"`        // 粉丝占比
		FemalePercent    float64 `json:"female_percent"`     // 女性占比
		MalePercent      float64 `json:"male_percent"`       // 男性占比
		Age1Percent      float64 `json:"age1_percent"`       // 年龄 <18
		Age2Percent      float64 `json:"age2_percent"`       // 年龄 18~24
		Age3Percent      float64 `json:"age3_percent"`       // 年龄 25~34
		Age4Percent      float64 `json:"age4_percent"`       // 年龄 35~44
		Age5Percent      float64 `json:"age5_percent"`       // 年龄 >44
		PhoneTop1        string  `json:"phone_top1"`         // 手机型号top3 top1
		PhoneTop1Rate    float64 `json:"phone_top1_rate"`    // 手机型号top3 top1占比
		PhoneTop2        string  `json:"phone_top2"`         // 手机型号top3 top2
		PhoneTop2Rate    float64 `json:"phone_top2_rate"`    // 手机型号top3 top2占比
		PhoneTop3        string  `json:"phone_top3"`         // 手机型号top3 top3
		PhoneTop3Rate    float64 `json:"phone_top3_rate"`    // 手机型号top3 top3占比
		ProvinceTop1     string  `json:"province_top1"`      // 地域分布top3 top1
		ProvinceTop1Rate float64 `json:"province_top1_rate"` // 地域分布top3 top1占比
		ProvinceTop2     string  `json:"province_top2"`      // 地域分布top3 top2
		ProvinceTop2Rate float64 `json:"province_top2_rate"` // 地域分布top3 top2
		ProvinceTop3     string  `json:"province_top3"`      // 地域分布top3 top3
		ProvinceTop3Rate float64 `json:"province_top3_rate"` // 地域分布top3 top3占比
		InterestTop1     string  `json:"interest_top1"`      // 兴趣分布top3 top1
		InterestTop1Rate float64 `json:"interest_top1_rate"` // 兴趣分布top3 top1占比
		InterestTop2     string  `json:"interest_top2"`      // 兴趣分布top3 top2
		InterestTop2Rate float64 `json:"interest_top2_rate"` // 兴趣分布top3 top2占比
		InterestTop3     string  `json:"interest_top3"`      // 兴趣分布top3 top3
		InterestTop3Rate float64 `json:"interest_top3_rate"` // 兴趣分布top3 top3占比
	}

	// CommentCompData 评论区组件效果数据
	CommentCompData struct {
		CompType       int32   `json:"comp_type"`         // 组件类型
		CompContent    string  `json:"comp_content"`      // 组件文案
		CompImpNum     int32   `json:"comp_imp_num"`      // 组件曝光量
		CompClickUvNum int32   `json:"comp_click_uv_num"` // 组件点击量
		CompClickPvNum int32   `json:"comp_click_pv_num"` // 组件点击人数
		ClickPvRate    float64 `json:"click_pv_rate"`     // 组件CTR
	}

	// ContentCompData 正文组件数据
	ContentCompData struct {
		WithContentComp       int32   `json:"with_content_comp"`        // 绑定正文组件 是否绑定正文组件，1-是，0-否
		ContentCompType       int32   `json:"content_comp_type"`        // 正文组件类型
		ContentCompContent    string  `json:"content_comp_content"`     // 正文组件文案
		ContentCompImpNum     int32   `json:"content_comp_imp_num"`     // 正文组件曝光量
		ContentCompClickNum   int32   `json:"content_comp_click_num"`   // 正文组件点击量
		ContentCompClickUv    int32   `json:"content_comp_click_uv"`    // 正文组件点击人数
		ConnentCompClickRatio float64 `json:"connent_comp_click_ratio"` // 正文组件点击率。如果50%，则返回50.00
	}
	// NoteBottomCompData 笔记底栏组件
	NoteBottomCompData struct {
		WithNoteBottomComp       int32   `json:"with_note_bottom_comp"`         // 是否使用笔记底栏组件，0-否，1-是
		NoteBottomCompType       int32   `json:"note_bottom_comp_type"`         // 组件类型 1-商品组件 2-店铺组件
		NoteBottomCompContent    string  `json:"note_bottom_comp_content"`      // 	笔记底栏组件文案
		NoteBottomCompImpNum     int32   `json:"note_bottom_comp_imp_num"`      // 	底栏曝光
		NoteBottomCompClickPvNum int32   `json:"note_bottom_comp_click_pv_num"` // 	底栏点击数
		NoteBottomCompClickUvNum int32   `json:"note_bottom_comp_click_uv_num"` // 	底栏点击人数
		NoteBottomCompClickRatio float64 `json:"note_bottom_comp_click_ratio"`  // 	底栏点击率
	}
	// EngageCompData 互动组件
	EngageCompData struct {
		WithEngageComp       int32   `json:"with_engage_comp"`         // 是否有互动组件 0-否 1-是
		EngageCompType       int32   `json:"engage_comp_type"`         // 组件类型 1-PK组件 2-投票组件
		EngageCompTitle      string  `json:"engage_comp_title"`        // 互动组件标题
		EngageCompImpNum     int32   `json:"engage_comp_imp_num"`      // 组件总曝光人数
		EngageCompClickUvNum int32   `json:"engage_comp_click_uv_num"` // 参与人数
		EngageCompClickRatio float64 `json:"engage_comp_click_ratio"`  // 参与率
	}

	// StarTaskData 星任务数据
	StarTaskData struct {
		// 基础数据
		StarTaskName        string  `json:"star_task_name"`         // 星任务名称	如果没数据，则为空串
		StarEventGroupId    string  `json:"star_event_group_id"`    // 星任务主任务Id	如果没数据，则为空串
		EventGroupStartTime string  `json:"event_group_start_time"` // 星任务考试时间
		EventGroupEndTime   string  `json:"event_group_end_time"`   // 星任务结束时间
		StarTotalAmount     float64 `json:"star_total_amount"`      // 总金额	单位：元
		StarPgyTotalAmount  float64 `json:"star_pgy_total_amount"`  // 蒲公英金额	单位：元
		StarAdsTotalAmount  float64 `json:"star_ads_total_amount"`  // 广告金额	单位：元
		StarTransRatio      float64 `json:"star_trans_ratio"`       // 抽样比例	传输比例为50%，则返回50.00
		// 累计数据
		StarReadUv          int32   `json:"star_read_uv"`           // 阅读UV
		StarCmtUv           int32   `json:"star_cmt_uv"`            // 评论UV
		StarLikeUv          int32   `json:"star_like_uv"`           // 点赞UV
		StarFavUv           int32   `json:"star_fav_uv"`            // 收藏UV
		StarShareUv         int32   `json:"star_share_uv"`          // 分享UV
		StarEnterStoreUv    int32   `json:"star_enter_store_uv"`    // 站外店铺行为UV
		StarEnterStoreRatio float64 `json:"star_enter_store_ratio"` // 站外转化率	站外转化率=（站外店铺行为UV/抽样比例）/阅读uv如转化率为50%，则返回50.00
		StarEnterStoreCost  float64 `json:"star_enter_store_cost"`  // 站外转化成本	站外转化成本=（该笔记订单成本/综合进店uv）*数据回传比例单位：元
	}
)
type NoteData struct {
	//下单账号信息
	DateKey         string `json:"date_key"`          // 更新日期
	OperateUserId   string `json:"operate_user_id"`   // 下单账号id
	OperateUserName string `json:"operate_user_name"` // 下单账号名称
	OperateMainName string `json:"operate_main_name"` // 下单账号主体名称
	BrandUserId     string `json:"brand_user_id"`     // 报备品牌id
	BrandUserName   string `json:"brand_user_name"`   // 报备品牌名称
	// 博主基础信息
	KolNickName    string `json:"kol_nick_name"`    // 博主昵称
	KolId          string `json:"kol_id"`           // 博主userId
	KolLink        string `json:"kol_link"`         // 博主主页链接
	KolFanNum      int    `json:"kol_fan_num"`      // 博主粉丝量
	KolCreditLevel int    `json:"kol_credit_level"` // 博主信用等级
	OrderKolLevel  int    `json:"order_kol_level"`  // 下单时博主信用等级
	// 笔记基础信息
	NoteTitle       string `json:"note_title"`        // 笔记标题
	NoteLink        string `json:"note_link"`         // 笔记链接
	NoteType        int    `json:"note_type"`         // 笔记类型
	NotePublishTime string `json:"note_publish_time"` // 笔记发布日期
	CooperateType   int    `json:"cooperate_type"`    // 笔记来源
	NoteId          string `json:"note_id"`           // 笔记id
	Duration        int32  `json:"duration"`          // 视频笔记视频总时长
	//订单基础信息
	OrderId            string `json:"order_id"`             // 订单id
	BizTitle           string `json:"biz_title	"`           // 订单标题
	KolPrice           int32  `json:"kol_price"`            // 博主报价
	TotalPlatformPrice int32  `json:"total_platform_price"` // 服务费金额
	Effect             bool   `json:"effect"`               // 是否为优效模式
	//全部流量效果
	ImpNum           int32   `json:"imp_num"`           //	曝光量
	HeatImpNum       int32   `json:"heat_imp_num"`      //	加热曝光量
	PromotionImpNum  int32   `json:"promotion_imp_num"` // 推广曝光量
	ReadNum          int32   `json:"read_num"`
	HeatReadNum      int32   `json:"heat_read_num"`      //	加热阅读量
	PromotionReadNum int32   `json:"promotion_read_num"` // 推广阅读量
	ReadUv           int32   `json:"read_uv"`            // 	阅读UV 备注：从xx开始有
	VideoPlay5SRate  float64 `json:"video_play_5s_rate"` // 5s播放率
	PicRead3SRate    float64 `json:"pic_read_3s_rate"`   // 3s阅读率
	AvgViewTime      int32   `json:"avg_view_time"`      // 平均浏览时长
	FinishRate       float64 `json:"finish_rate"`        // 视频完播率
	// 全部互动效果
	EngageNum  int32   `json:"engage_num"`  // 互动量
	EngageRate float64 `json:"engage_rate"` // 互动率
	LikeNum    int32   `json:"like_num"`    // 点赞量
	FavNum     int32   `json:"fav_num"`     // 收藏量
	CmtNum     int32   `json:"cmt_num"`     // 评论量
	ShareNum   int32   `json:"share_num"`   // 分享量
	// 自然流量效果
	OriginImpNum  int32 `json:"origin_imp_num"`  // 曝光量
	OriginReadNum int32 `json:"origin_read_num"` // 阅读量
	// 曝光来源分布占比（自然流量）
	ImpDiscovery float64 `json:"imp_discovery"` // 发现页
	ImpSearch    float64 `json:"imp_search"`    // 搜索页
	ImpHomepage  float64 `json:"imp_homepage"`  // 个人页
	ImpFollow    float64 `json:"imp_follow"`    // 关注页
	ImpNearby    float64 `json:"imp_nearby"`    // 附近页
	ImpOther     float64 `json:"imp_other"`     // 其他
	// 阅读来源分布占比（自然流量）
	ReadDiscovery float64 `json:"read_discovery"` // 发现页
	ReadSearch    float64 `json:"read_search"`    // 搜索页
	ReadHomepage  float64 `json:"read_homepage"`  // 个人页
	ReadFollow    float64 `json:"read_follow"`    // 关注页
	ReadNearby    float64 `json:"read_nearby"`    // 附近页
	ReadOther     float64 `json:"read_other"`     // 其他
	// 性价比情况
	ReadCost   int32 `json:"read_cost"`   // 阅读单价
	EngageCost int32 `json:"engage_cost"` // 互动单价
	Cpcp       int32 `json:"cpcp"`        // 消费意向单价
	// 评论区组件效果数据
	// CompType       int32   `json:"comp_type"`         // 组件类型
	// CompContent    string  `json:"comp_content"`      // 组件文案
	// CompImpNum     int32   `json:"comp_imp_num"`      // 组件曝光量
	// CompClickUvNum int32   `json:"comp_click_uv_num"` // 组件点击量
	// CompClickPvNum int32   `json:"comp_click_pv_num"` // 组件点击人数
	// ClickPvRate    float64 `json:"click_pv_rate"`     // 组件CTR
	CommentCompData
	//正文组件数据
	// WithContentComp       int32   `json:"with_content_comp"`        // 绑定正文组件 是否绑定正文组件，1-是，0-否
	// ContentCompType       int32   `json:"content_comp_type"`        // 正文组件类型
	// ContentCompContent    string  `json:"content_comp_content"`     // 正文组件文案
	// ContentCompImpNum     int32   `json:"content_comp_imp_num"`     // 正文组件曝光量
	// ContentCompClickNum   int32   `json:"content_comp_click_num"`   // 正文组件点击量
	// ContentCompClickUv    int32   `json:"content_comp_click_uv"`    // 正文组件点击人数
	// ConnentCompClickRatio float64 `json:"connent_comp_click_ratio"` // 正文组件点击率。如果50%，则返回50.00
	ContentCompData
	// 消费意向
	Cp     int     `json:"cp"`      // 消费意向 如果没有spu信息，则为0
	CpRate float64 `json:"cp_rate"` // 消费意向转化率 如果没有spu信息，则为0
	// 阅读受众分析
	AudiencePersonaData
	// 星任务数据
	// // 基础数据
	// StarTaskName        string  `json:"star_task_name"`         // 星任务名称	如果没数据，则为空串
	// StarEventGroupId    string  `json:"star_event_group_id"`    // 星任务主任务Id	如果没数据，则为空串
	// EventGroupStartTime string  `json:"event_group_start_time"` // 星任务考试时间
	// EventGroupEndTime   string  `json:"event_group_end_time"`   // 星任务结束时间
	// StarTotalAmount     float64 `json:"star_total_amount"`      // 总金额	单位：元
	// StarPgyTotalAmount  float64 `json:"star_pgy_total_amount"`  // 蒲公英金额	单位：元
	// StarAdsTotalAmount  float64 `json:"star_ads_total_amount"`  // 广告金额	单位：元
	// StarTransRatio      float64 `json:"star_trans_ratio"`       // 抽样比例	传输比例为50%，则返回50.00
	// // 累计数据
	// StarReadUv          int32   `json:"star_read_uv"`           // 阅读UV
	// StarCmtUv           int32   `json:"star_cmt_uv"`            // 评论UV
	// StarLikeUv          int32   `json:"star_like_uv"`           // 点赞UV
	// StarFavUv           int32   `json:"star_fav_uv"`            // 收藏UV
	// StarShareUv         int32   `json:"star_share_uv"`          // 分享UV
	// StarEnterStoreUv    int32   `json:"star_enter_store_uv"`    // 站外店铺行为UV
	// StarEnterStoreRatio float64 `json:"star_enter_store_ratio"` // 站外转化率	站外转化率=（站外店铺行为UV/抽样比例）/阅读uv如转化率为50%，则返回50.00
	// StarEnterStoreCost  float64 `json:"star_enter_store_cost"`  // 站外转化成本	站外转化成本=（该笔记订单成本/综合进店uv）*数据回传比例单位：元
	StarTaskData
	// 种草数据
	SpuName string `json:"spu_name"` // SPU名称
	// 种草指标
	InterestNum       int32   `json:"interest_num"`        // 种草数
	InterestRate      float64 `json:"interest_rate"`       // 种草率	=种草值/阅读量。如果50%，则返回50.00
	InterestCost      float64 `json:"interest_cost"`       // 种草成本	=蒲公英笔记消耗/全部流量下的种草值。单位：元
	FeedInterestNum   int32   `json:"feed_interest_num"`   // 推荐场种草数
	SearchInterestNum int32   `json:"search_interest_num"` // 搜索场种草数
	OtherInterestNum  int32   `json:"other_interest_num"`  // 其他场种草数
	// 笔记底栏组件
	// WithNoteBottomComp       int32   `json:"with_note_bottom_comp"`         // 是否使用笔记底栏组件，0-否，1-是
	// NoteBottomCompType       int32   `json:"note_bottom_comp_type"`         // 组件类型 1-商品组件 2-店铺组件
	// NoteBottomCompContent    string  `json:"note_bottom_comp_content"`      // 	笔记底栏组件文案
	// NoteBottomCompImpNum     int32   `json:"note_bottom_comp_imp_num"`      // 	底栏曝光
	// NoteBottomCompClickPvNum int32   `json:"note_bottom_comp_click_pv_num"` // 	底栏点击数
	// NoteBottomCompClickUvNum int32   `json:"note_bottom_comp_click_uv_num"` // 	底栏点击人数
	// NoteBottomCompClickRatio float64 `json:"note_bottom_comp_click_ratio"`  // 	底栏点击率
	NoteBottomCompData
	// 互动组件
	// WithEngageComp       int32   `json:"with_engage_comp"`         // 是否有互动组件 0-否 1-是
	// EngageCompType       int32   `json:"engage_comp_type"`         // 组件类型 1-PK组件 2-投票组件
	// EngageCompTitle      string  `json:"engage_comp_title"`        // 互动组件标题
	// EngageCompImpNum     int32   `json:"engage_comp_imp_num"`      // 组件总曝光人数
	// EngageCompClickUvNum int32   `json:"engage_comp_click_uv_num"` // 参与人数
	// EngageCompClickRatio float64 `json:"engage_comp_click_ratio"`  // 参与率
	EngageCompData
}

type NotePostData struct {
	Total     int        `json:"total"`
	PageNum   int        `json:"page_num"`
	PageSize  int        `json:"page_size"`
	TotalPage int        `json:"total_page"`
	Items     []NoteData `json:"datas"`
}

type NoteDataRequest struct {
	UserId    string `json:"user_id"`    //报备品牌账号id，鉴权使用
	DateType  int    `json:"date_type"`  //日期类型，1为订单创建时间，2为笔记发布时间
	StartTime string `json:"start_time"` //格式yyyy-MM-ddstart_time和end_time的时间跨度不能超过30天如设置start_time=2022-12-01，end_time=2022-12-31则查询 2022-12-01 00:00:00 ~ 2022-12-31 23:59:59 之间的数据
	EndTime   string `json:"end_time"`   //格式yyyy-MM-dd
	PageNum   int    `json:"page_num"`   //当前第几页
	PageSize  int    `json:"page_size"`  //每页多少条，最大支持100
	DateKey   string `json:"date_key"`   //日期（格式yyyyMMdd），获取最新数据时不传该参数，可以支持历史三天的数据拉取，小红星数据为T+2，其他数据可以支持T+1
}

type ListNotePostResponse struct {
	ApiResp
	Data NotePostData `json:"data"`
}

func (s *NoteService) ListNotePostData(ctx context.Context, req *NoteDataRequest, options ...RequestOption) (*ListNotePostResponse, error) {
	path := "/api/open/pgy/note/post/data"

	// apiReq.Body = map[string]interface{}{"user_id": userId, "date_type": 2, "start_time": "2024-08-01", "end_time": "2024-08-16", "page_num": 1, "page_size": 100}
	// body = req

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &ListNotePostResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
