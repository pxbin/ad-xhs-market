package adxhsmarket

// ReportService 表示聚光数据报名服务
type ReportService service

// DataReportDTO 数据指标
type DataReportDTO struct {
	// 基础指标
	Fee        float64 `json:"fee,string"`        // 消费	 推广消费金额 单位:元
	Impression int64   `json:"impression,string"` // 展现量	推广展现量
	Click      int64   `json:"click,string"`      // 点击量 推广点击量；视频流中「观看视频5s」记做一次点击
	Ctr        float64 `json:"ctr,string"`        // 点击率	推广点击量/推广展现量*100%
	Acp        float64 `json:"acp,string"`        // 平均点击成本	推广消费/推广点击量
	Cpm        float64 `json:"cpm,string"`        // 平均千次曝光	推广消费/推广展现量*1000
	// 互动指标
	Like              int64   `json:"like,string"`                // 点赞	用户点击推广后1小时内，对笔记的点赞量
	Comment           int64   `json:"comment,string"`             // 评论	用户点击推广后1小时内，对笔记的评论量
	Collect           int64   `json:"collect,string"`             // 收藏	用户点击推广后1小时内，对笔记的收藏量
	Follow            int64   `json:"follow,string"`              // 关注	用户点击推广后，对笔记作者的关注量
	Share             int64   `json:"share,string"`               // 分享	用户点击推广后1小时内，对笔记进行站内或站外分享量, 支持2021.07.07之后的数据
	Interaction       int64   `json:"interaction,string"`         // 互动量	点赞、收藏、关注、评论、分享5个互动行为指标之和
	Cpi               float64 `json:"cpi,string"`                 // 平均互动成本	消费/互动量
	ActionButtonClick int64   `json:"action_button_click,string"` // 行动按钮点击量	不同类型创意包含的组件点击量，包含预告组件点击
	ActionButtonCtr   float64 `json:"action_button_ctr,string"`   // 行动按钮点击率	行动按钮点击量/点击量
	Screenshot        int64   `json:"screenshot,string"`          // 截图	用户点击推广后1小时内，对笔记进行截图的次数
	PicSave           int64   `json:"pic_save,string"`            // 保存图片	用户点击推广后1小时内，保存笔记中图片的次数
	ReservePv         int64   `json:"reserve_pv,string"`          // 预告组件点击	用户点击推广1小时内，在账户内的预约人次。不去重
	// 直播间互动
	ClkLiveEntryPv       int64   `json:"clk_live_entry_pv,string"`         // 直播间观看次数	用户点击推广后24小时内，进入直播间的总次数
	ClkLiveEntryPvCost   float64 `json:"clk_live_entry_pv_cost,string"`    // 直播间观看成本	推广消费/直播间观看次数
	ClkLiveAvgViewTime   int64   `json:"clk_live_avg_view_time,string"`    // 播间人均停留时长	用户点击推广后24小时内，进入直播间后的平均观看时长（单位：分钟）
	ClkLiveAllFollow     int64   `json:"clk_live_all_follow,string"`       // 直播间新增粉丝量	用户点击推广后24小时内，进入直播间后在直播间内关注主播的总次数
	ClkLive5SEntryPv     int64   `json:"clk_live_5s_entry_pv,string"`      // 直播间有效观看次数	用户点击推广后24小时内，进入直播间后停留时长大于等于5秒的总次数
	ClkLive5SEntryUvCost float64 `json:"clk_live_5s_entry_uv_cost,string"` // 直播间有效观看成本	推广消费/直播间有效观看人次
	ClkLiveComment       int64   `json:"clk_live_comment,string"`          // 直播间评论次数	用户点击推广后24小时内，进入直播间后在直播间内评论的总次数
	// 笔记种草指标
	SearchCmtClick        int64   `json:"search_cmt_click,string"`          // 搜索组件点击量	点击推广搜索组件的次数
	SearchCmtClickCvr     float64 `json:"search_cmt_click_cvr,string"`      // 搜索组件点击转化率	搜索组件点击量/点击量
	SearchCmtAfterRead    int64   `json:"search_cmt_after_read,string"`     // 搜后阅读量	点击组件后的搜索场景阅读量
	SearchCmtAfterReadAvg float64 `json:"search_cmt_after_read_avg,string"` // 平均搜后阅读笔记篇数	搜后阅读量/搜索组件点击量
	// 电商转化指标/购买兴趣
	GoodsVisit       int64   `json:"goods_visit,omitempty,string"`        // 进店访问量	用户点击推广后的7日内，对客户所绑定小红书店铺首页或商品详情页的访问总次数
	GoodsVisitPrice  float64 `json:"goods_visit_price,omitempty,string"`  // 进店访问成本	推广消费/进店访问量
	SellerVisit      int64   `json:"seller_visit,omitempty,string"`       // 商品访客量	用户点击推广后的7日内，对客户所绑定小红书店铺首页或商品详情页的访问总人数
	SellerVisitPrice float64 `json:"seller_visit_price,omitempty,string"` // 商品访客成本	推广消费/商品访客量
	ShoppingCartAdd  int64   `json:"shopping_cart_add,omitempty,string"`  // 商品加购量	用户点击推广后的7日内，对客户所绑定小红书店铺全部商品加入购物车的总次数
	AddCartPrice     float64 `json:"add_cart_price,omitempty,string"`     // 商品加购成本	推广消费/商品加购量
	// 电商转化指标/7日转化
	PresaleOrderNum7D    string `json:"presale_order_num_7d"`    // 7日预售订单量	用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部预售商品，完成预售定金支付的总订单数
	PresaleOrderGmv7D    string `json:"presale_order_gmv_7d"`    // 7日预售订单金额	用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部预售商品，完成预售定金支付的订单所对应的商品总金额（含定金与应支付尾款）
	GoodsOrder           string `json:"goods_order"`             // 7日下单订单量	用户点击广告后，7天内在广告主绑定的店铺，对全部商品的创建订单数之和
	GoodsOrderPrice      string `json:"goods_order_price"`       // 7日下单订单成本	推广消费/7日下单订单量
	Rgmv                 string `json:"rgmv"`                    // 7日下单金额	用户点击广告后，7天内在广告主绑定的店铺，对全部商品的支付金额之和（含退款金额）
	Roi                  string `json:"roi"`                     // 7日下单ROI	rgmv/消费
	SuccessGoodsOrder    string `json:"success_goods_order"`     // 7日支付订单量	用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部商品的支付总订单数量
	ClickOrderCvr        string `json:"click_order_cvr"`         // 7日支付转化率	7日支付订单量/推广点击数*100%
	PurchaseOrderPrice7D string `json:"purchase_order_price_7d"` // 7日支付金额	用户点击推广后的7日内，对客户所绑定小红书店铺和直播间全部商品的支付总订单金额
	PurchaseOrderGmv7D   string `json:"purchase_order_gmv_7d"`   // 7日支付ROI	7日支付金额/推广消费
	PurchaseOrderRoi7D   string `json:"purchase_order_roi_7d"`   // 7日支付ROI	7日支付金额/推广消费
	// 电商转化指标/直播间转化
	ClkLiveRoomOrderNum  string `json:"clk_live_room_order_num"` // 直播间支付订单量	用户点击推广后的24小时内，在直播间内全部商品的支付总订单量
	LiveAverageOrderCost string `json:"live_average_order_cost"` // 直播间支付订单成本	推广消费/直播间支付订单量
	ClkLiveRoomRgmv      string `json:"clk_live_room_rgmv"`      // 直播间支付金额	用户点击推广后的24小时内，在直播间内全部商品的支付总订单金额
	ClkLiveRoomRoi       string `json:"clk_live_room_roi"`       // 直播间支付ROI	直播间支付金额/推广消费
	// 销售线索指标
	Leads                 string `json:"leads"`                   // 表单提交	用户点击广告后成功提交表单的次数
	LeadsCpl              string `json:"leads_cpl"`               // 表单成本	消费/表单数
	LandingPageVisit      string `json:"landing_page_visit"`      // 落地页访问量（行为时间）	推广落地页访问量，时间记录在行为时间上
	LeadsButtonImpression string `json:"leads_button_impression"` // 表单按钮曝光量（行为时间）	推广表单按钮曝光量，时间记录在行为时间上
	ValidLeads            string `json:"valid_leads"`             // 有效表单	用户提交表单后7天内，有过进一步的沟通行为（需要您回传此事件）；转化记录在计费时间上
	ValidLeadsCpl         string `json:"valid_leads_cpl"`         // 有效表单成本	消费/有效表单
	LeadsCvr              string `json:"leads_cvr"`               // 表单转化率	表单数/点击数
	PhoneCallCnt          string `json:"phone_call_cnt"`          // 电话拨打	用户点击推广后拨打电话的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	PhoneCallSuccCnt      string `json:"phone_call_succ_cnt"`     // 电话接通	用户拨打电话后7天内，能拨通电话的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	WechatCopyCnt         string `json:"wechat_copy_cnt"`         // 微信复制	用户点击推广后，用户点击推广后成功复制微信号的次数，（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	WechatCopySuccCnt     string `json:"wechat_copy_succ_cnt"`    // 微信加为好友	用户复制微信后7天内，成功复制微信且成功加为好友（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	IdentityCertiCnt      string `json:"identity_certi_cnt"`      // 身份认证	用户点击推广后，用户点击推广后成功完成身份认证的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	CommodityBuyCnt       string `json:"commodity_buy_cnt"`       // 商品购买	用户点击推广后，用户点击推广后成功完成付费的次数（仅支持自研落地页，需要您回传此事件）；转化记录在计费时间上
	// 私信营销指标
	MessageUser              string `json:"message_user"`           // 私信咨询人数	用户点击推广后24小时内，通过推广产生私信对话的用户数（包括点击私信组件、专业号主页咨询等）
	Message                  string `json:"message"`                // 私信咨询条数	用户点击推广后24小时内，通过推广产生私信对话的条数
	MessageConsult           string `json:"message_consult"`        // 私信咨询数	用户点击推广后24小时内，至少产生过一次咨询记一个私信咨询数（包括点击私信组件、专业号主页咨询等）
	InitiativeMessage        string `json:"initiative_message"`     // 私信开口数	用户点击推广24小时内，至少产生过一次开口行为记一个私信开口数
	MessageConsultCpl        string `json:"message_consult_cpl"`    // 私信咨询成本	消耗/私信咨询数
	InitiativeMessageCpl     string `json:"initiative_message_cpl"` // 私信开口成本	消耗/私信开口数
	MsgLeadsNum              string `json:"msg_leads_num"`          // 私信留资数	用户点击推广后7日内，成功留资转化的次数（需要在私信回复工具上进行客资标注）
	MsgLeadsCost             string `json:"msg_leads_cost"`         // 私信留资成本	消耗/私信留资数
	MessageFstReplyTimeAvg   string `json:"message_fst_reply_time_avg"`
	MessageReplyIn3MinRate   string `json:"message_reply_in3_min_rate"`
	FstMessageReplyIn45SRate string `json:"fst_message_reply_in45s_rate"`
	// 行业商品销量指标
	ExternalGoodsVisit24H      string `json:"external_goods_visit_24h"`       // 行业商品点击量	用户点击广告组件进行跳转的行为次数
	ExternalGoodsVisitPrice24H string `json:"external_goods_visit_price_24h"` // 行业商品点击成本	消费 / 行业和商品点击量
	ExternalGoodsVisitRate24H  string `json:"external_goods_visit_rate_24h"`  // 行业商品点击转化率	行业商品点击量 / 点击量
	ExternalGoodsVisit7        string `json:"external_goods_visit_7"`         // 行业商品点击量	用户点击广告组件进行跳转的行为次数（字段取值变更，原external_goods_visit_24h字段）
	ExternalGoodsVisitPrice7   string `json:"external_goods_visit_price_7"`   // 行业商品点击成本	消费 / 行业和商品点击量（字段取值变更，原 external_goods_visit_price_24h字段）
	ExternalGoodsVisitRate7    string `json:"external_goods_visit_rate_7"`    // 行业商品点击转化率	行业商品点击量 / 点击量（字段取值变更，原external_goods_visit_rate_24h字段）
	ExternalGoodsOrder7        string `json:"external_goods_order_7"`         // 行业商品成交订单量（7日)	用户点击广告后，7天内对全部商品的创建且成功支付订单数之和。视频流场景下，用户曝光广告后，7天内对全部商品的创建且成功支付订单数之和。
	ExternalRgmv7              string `json:"external_rgmv_7"`                // 行业商品GMV（7日）	用户点击广告后，7天内对全部商品的成交金额之和。视频流场景下，用户曝光广告后，7天内对全部商品的成交金额之和。
	ExternalGoodsOrderPrice7   string `json:"external_goods_order_price_7"`   // 行业商品成交订单成本（7日）	消费 / 行业商品成交订单量(7日)
	ExternalGoodsOrderRate7    string `json:"external_goods_order_rate_7"`    // 行业商品成交订单转化率（7日）
	ExternalRoi7               string `json:"external_roi_7"`                 // 行业商品ROI（7日）	行业商品GMV(7日) / 消费
	ExternalGoodsOrder15       string `json:"external_goods_order_15"`        // 行业商品成交订单量（15日）	用户点击广告后，15天内对全部商品的创建且成功支付订单数之和。视频流场景下，用户曝光广告后，15天内对全部商品的创建且成功支付订单数之和。
	ExternalRgmv15             string `json:"external_rgmv_15"`               // 行业商品GMV（15日）	用户点击广告后，15天内对全部商品的成交金额之和。视频流场景下，用户曝光广告后，15天内对全部商品的成交金额之和。
	ExternalGoodsOrderPrice15  string `json:"external_goods_order_price_15"`  // 行业商品成交订单成本（15日）	消费 / 行业商品成交订单量(15日)
	ExternalGoodsOrderRate15   string `json:"external_goods_order_rate_15"`   // 行业商品成交订单转化率（15日）	行业商品成交订单量(15日) / 点击量
	ExternalRoi15              string `json:"external_roi_15"`                // 行业商品ROI（15日）	行业商品GMV(15日) / 消费
	ExternalGoodsOrder30       string `json:"external_goods_order_30"`        // 行业商品成交订单量（30日）	用户点击广告后，30天内对全部商品的创建且成功支付订单数之和。视频流场景下，用户曝光广告后，30天内对全部商品的创建且成功支付订单数之和。
	ExternalRgmv30             string `json:"external_rgmv_30"`               // 行业商品GMV（30日）	用户点击广告后，30天内对全部商品的成交金额之和。视频流场景下，用户曝光广告后，30天内对全部商品的成交金额之和。
	ExternalGoodsOrderPrice30  string `json:"external_goods_order_price_30"`  // 行业商品成交订单成本（30日）	消费 / 行业商品成交订单量(30日)
	ExternalGoodsOrderRate30   string `json:"external_goods_order_rate_30"`   // 行业商品成交订单转化率（30日）	行业商品成交订单量(30日) / 点击量
	ExternalRoi30              string `json:"external_roi_30"`                // 行业商品ROI（30日）	行业商品GMV(30日) / 消费
	// 外链专属指标
	ExternalLeads    string `json:"external_leads"`     // 外链转化数	外链回传的转化数据，例如：表单数、下载数、付费数
	ExternalLeadsCpl string `json:"external_leads_cpl"` // 平均外链转化成本	消费/转化数
	// 关键词指标
	WordAvgLocation         string `json:"word_avg_location"`          // 平均位次
	WordImpressionRankFirst string `json:"word_impression_rank_first"` // 首位曝光排名
	WordImpressionRateFirst string `json:"word_impression_rate_first"` // 首位曝光占比
	WordImpressionRankThird string `json:"word_impression_rank_third"` // 前三位曝光排名
	WordImpressionRateThird string `json:"word_impression_rate_third"` // 前三位曝光占比
	WordClickRankFirst      string `json:"word_click_rank_first"`      // 首位点击排名
	WordClickRateFirst      string `json:"word_click_rate_first"`      // 首位点击占比
	WordClickRateThird      string `json:"word_click_rate_third"`      // 前三位点击占比
	WordClickRankThird      string `json:"word_click_rank_third"`      // 前三位点击排名
	// 唤端指标
	InvokeAppOpenCnt            string `json:"invoke_app_open_cnt"`             // APP打开量（唤起）
	InvokeAppOpenCost           string `json:"invoke_app_open_cost"`            // APP打开成本（唤起）
	InvokeAppEnterStoreCnt      string `json:"invoke_app_enter_store_cnt"`      // APP进店量（唤起）
	InvokeAppEnterStoreCost     string `json:"invoke_app_enter_store_cost"`     // APP进店成本（唤起）
	InvokeAppEngagementCnt      string `json:"invoke_app_engagement_cnt"`       // APP互动量（唤起）
	InvokeAppEngagementCost     string `json:"invoke_app_engagement_cost"`      // APP互动成本（唤起）
	InvokeAppPaymentCnt         string `json:"invoke_app_payment_cnt"`          // APP支付次数（唤起）
	InvokeAppPaymentCost        string `json:"invoke_app_payment_cost"`         // APP订单支付成本（唤起）
	SearchInvokeButtonClickCnt  string `json:"search_invoke_button_click_cnt"`  // APP打开按钮点击量（唤起）
	SearchInvokeButtonClickCost string `json:"search_invoke_button_click_cost"` // APP打开按钮点击成本（唤起
	// 京东站外店铺行为指标
	JdActiveUserNum    string `json:"jd_active_user_num"`     // 京东站外活跃行为量
	JdActiveUserNumCvr string `json:"jd_active_user_num_cvr"` // 京东站外转化率
	JdActiveUserNumCpl string `json:"jd_active_user_num_cpl"` // 京东站外转化成本
	// 小红星站外店铺行为指标
	OutsideShopVisit      string `json:"outside_shop_visit"`       // 小红星站外进店量
	OutsideShopVisitPrice string `json:"outside_shop_visit_price"` // 小红星站外进店成本
	OutsideShopVisitRate  string `json:"outside_shop_visit_rate"`  // 小红星站外进店率
	// 应用下载指标
	AppDownloadButtonClickCnt  string `json:"app_download_button_click_cnt"`  // APP下载按钮点击
	AppDownloadButtonClickCtr  string `json:"app_download_button_click_ctr"`  // APP下载按钮点击率
	AppDownloadButtonClickCost string `json:"app_download_button_click_cost"` // APP下载按钮点击成本
	AppActivateCnt             string `json:"app_activate_cnt"`               // 激活数
	AppActivateCost            string `json:"app_activate_cost"`              // 激活成本
	AppActivateCtr             string `json:"app_activate_ctr"`               // 激活率
	AppRegisterCnt             string `json:"app_register_cnt"`               // 注册数
	AppRegisterCost            string `json:"app_register_cost"`              // 注册成本
	AppRegisterCtr             string `json:"app_register_ctr"`               // 注册率
	FirstAppPayCnt             string `json:"first_app_pay_cnt"`              // 首次付费数
	FirstAppPayCost            string `json:"first_app_pay_cost"`             // 首次付费成本
	FirstAppPayCtr             string `json:"first_app_pay_ctr"`              // 首次付费率
	CurrentAppPayCnt           string `json:"current_app_pay_cnt"`            // 当日付费次数
	CurrentAppPayCost          string `json:"current_app_pay_cost"`           // 当日付费成本
	AppKeyActionCnt            string `json:"app_key_action_cnt"`             // 关键行为数
	AppKeyActionCost           string `json:"app_key_action_cost"`            // 关键行为成本
	AppKeyActionCtr            string `json:"app_key_action_ctr"`             // 关键行为率
	AppPayCnt7d                string `json:"app_pay_cnt_7d"`                 // 7日付费次数
	AppPayCost7d               string `json:"app_pay_cost_7d"`                // 7日付费成本
	AppPayAmount               string `json:"app_pay_amount"`                 // 付费金额
	AppPayRoi                  string `json:"app_pay_roi"`                    // 付费ROI
	AppActivateAmount1d        string `json:"app_activate_amount_1d"`         // 当日LTV
	AppActivateAmount3d        string `json:"app_activate_amount_3d"`         // 三日LTV
	AppActivateAmount7d        string `json:"app_activate_amount_7d"`         // 七日LTV
	AppActivateAmount1dRoi     string `json:"app_activate_amount_1d_roi"`     // 当日广告付费ROI
	AppActivateAmount3dRoi     string `json:"app_activate_amount_3d_roi"`     // 三日广告付费ROI
	AppActivateAmount7dRoi     string `json:"app_activate_amount_7d_roi"`     // 七日广告付费ROI
	Retention1dCnt             string `json:"retention_1d_cnt"`               // 次留
	Retention3dCnt             string `json:"retention_3d_cnt"`               // 3日留存
	Retention7dCnt             string `json:"retention_7d_cnt"`               // 7日留存
}

type (
	// BaseCreativityDTO 创意属性信息
	BaseCreativityDTO struct {
		CreativityId          int    `json:"creativity_id"`           // 创意id
		CreativityName        string `json:"creativity_name"`         // 创意名称
		CreativityFilterState int    `json:"creativity_filter_state"` // 创意状态8：有效，3：暂停，9: 商品状态异常，4：已被单元暂停，10：单元未开始，11：单元已结束，12：单元处于暂停时段，5：已被计划暂停，13：计划预算不足，16：账户日预算不足，14：现金余额不足，1：已删除
		CreativityCreateTime  string `json:"creativity_create_time"`  // 创意创建时间：格式 yyyy-MM-dd HH:mm:ss
		CreativityEnable      int    `json:"creativity_enable"`       // 创意启停状态：0：暂停，1：开启
		AuditStatus           int    `json:"audit_status"`            // 审核状态：0创建待审核，1审核通过，2审核拒绝，3更新待审核
		UnitId                int    `json:"unit_id"`                 // 单元id
		NoteId                string `json:"note_id"`                 // 笔记id
		Programmatic          int    `json:"programmatic"`            // 创意组合类型：0：自定义创意，1：程序化创意
		CreativityType        int    `json:"creativity_type"`         // 创意类型：0：笔记-无组件1：笔记-商品组件-购买同款商品2：笔记-商品组件-进店看看3：笔记-商品组件-小程序购买同款商品4：笔记-商品组件-小程序购买同款商品5：笔记-落地页组件-表单6：笔记-落地页组件-外跳链接7：笔记-私信组件8：笔记-直播间组件9：笔记-poi门店组件10：笔记-外链商品11：直播间12：搜索组件13：小程序组件14：留资组件
	}

	// BaseCampaignDTO 计划属性信息
	BaseCampaignDTO struct {
		CampaignId              int    `json:"campaign_id"`                // 计划id
		CampaignName            string `json:"campaign_name"`              // 计划名称
		CampaignFilterState     int    `json:"campaign_filter_state"`      // 计划状态
		CampaignCreateTime      string `json:"campaign_create_time"`       // 计划创建时间: 格式yyyy-MM-dd HH:mm:ss
		CampaignEnable          int    `json:"campaign_enable"`            // 计划启停状态：0：暂停，1：开启
		MarketingTarget         int    `json:"marketing_target"`           // 营销诉求:3：商品销量_日常推广，4：产品种草，8：直播推广_日常推广，9：客资收集，10：抢占赛道，14：直播推广_直播预热，15：商品销量_店铺拉新
		Placement               int    `json:"placement"`                  // 广告类型:1：信息流，2：搜索，4：全站智投，7：视频内流
		OptimizeTarget          int    `json:"optimize_target"`            // 推广目标：0：点击量1：互动量3：表单提交量4：商品成单量5：私信咨询量6：直播间观看量11：商品访客量12：落地页访问量13：私信开口量14：有效观看量18：站外转化量20：TI人群规模21：行业商品成单23：直播预热量24：直播间成交25：直播间支付ROI
		PromotionTarget         int    `json:"promotion_target"`           // 投放标的:1：笔记，2：商品，7：外链落地页，9：落地页，18：直播间
		BiddingStrategy         int    `json:"bidding_strategy"`           // 出价方式：2：手动出价3：自动出价
		ConstraintType          int    `json:"constraint_type"`            // 成本控制方式: -1: 无，101: 自动控制，0: 点击成本控制，1: 互动成本控制，3: 表单提交成本控制，5: 私信咨询成本控制，11: 访客成本控制，13: 私信开口成本控制，14: 有效观播成本控制，17: ROI控制，23: 预热成本控制，50: 私信留资成本控制
		ConstraintValue         int    `json:"constraint_value"`           // 成本控制值
		LimitDayBudget          int    `json:"limit_day_budget"`           //预算类型: 预算类型：0：不限预算，1：指定预算
		OriginCampaignDayBudget int    `json:"origin_campaign_day_budget"` // 计划日预算
		BudgetState             int    `json:"budget_state"`               // 预算状态，0: 计划预算不足，1 计划预算充足
		SmartSwitch             int    `json:"smart_switch"`               // 是否节假日预算上调，0: 关闭，1: 开启
		PacingMode              int    `json:"pacing_mode"`                // 投放速率，1: 匀速投放，2: 加速投放
		StartTime               string `json:"start_time"`                 // 计划开始时间：格式yyyy-MM-dd
		ExpireTime              string `json:"expire_time"`                // 计划结束时间：格式yyyy-MM-dd
		TimePeriod              string `json:"time_period"`                // 时段: 默认168个1：表示一周每个小时用0和1表示，0表示不投，1表示投放，示例中表示1点不投，其他时间投
		TimePeriodType          int    `json:"time_period_type"`           // 推广时段类型, 0: 全时段，1:自定义时间段
		BuildType               int    `json:"build_type"`                 // 搭建方式，0：标准搭建，1：省心智投
		SearchFlag              int    `json:"search_flag"`                // 是否信息流快投搜索：0: 否，1：是
		FeedFlag                int    `json:"feed_flag"`                  // 是否搜索追投信息流：0: 否，1：是
		MigrationStatus         int    `json:"migration_status"`           // 专业号平台计划迁移状态: 0：非迁移计划，2：迁移计划
	}

	// BaseUnitDTO  单元属性信息
	BaseUnitDTO struct {
		UnitId          int    `json:"unit_id"`           // 单元id
		UnitName        string `json:"unit_name"`         // 单元名称
		UnitFilterState int    `json:"unit_filter_state"` // 单元状态：10：有效，4：暂停，2：未开始，3： 已结束，5：处于暂停时段，6：已被计划暂停，8：计划预算不足，11：账户日预算不足，7：现金余额不足，1：已删除
		UnitCreateTime  string `json:"unit_create_time"`  // 单元创建时间：格式 yyyy-MM-dd HH:mm:ss
		UnitEnable      int    `json:"unit_enable"`       // 单元启停状态：0：暂停，1：开启
		CampaignId      int    `json:"campaign_id"`       // 计划id
		EventBid        int    `json:"event_bid"`         // 出价
	}

	// BaseTargetDTO 定向属性信息
	BaseTargetDTO struct {
		TargetName   string `json:"target_name"`   // 定向名称
		TargetStatus int    `json:"target_status"` // 定向状态
		UnitId       int64  `json:"unit_id"`       // 单元id
		CampaignId   int64  `json:"campaign_id"`   // 计划id
		TargetId     int64  `json:"target_id"`     // 定向id
	}

	// BaseKeywordDTO 表示关键词属性信息
	BaseKeywordDTO struct {
		KeywordId          int64  `json:"keyword_id"`           //	关键词id
		Keyword            string `json:"keyword"`              //	关键词
		UseBidStrategy     int    `json:"use_bid_strategy"`     //	出价策略：0：未使用出价策略1：已使用出价策略
		KeywordEnable      int    `json:"keyword_enable"`       //	关键词状态：0：未上线1：已上线
		KeywordFilterState int    `json:"keyword_filter_state"` //	关键词状态过滤2：删除3：暂停4：已被单元暂停5：已被计划暂停6：现金余额不足7：计划预算不足8：有效9：计划未开始10：计划已结束11：计划处于暂停阶段12：账户日预算不足
		UnitId             int64  `json:"unit_id"`              //	单元id
		CampaignId         string `json:"campaign_id"`          //	计划id
	}
)

type (
	// OfflineBizField 离线报表通用业务字段
	OfflineBizField struct {
		Time            string `json:"time"`                       // 业务字段	时间
		Placement       string `json:"placement,omitempty"`        // 业务字段	广告类型
		OptimizeTarget  string `json:"optimize_target,omitempty"`  // 业务字段	优化目标
		PromotionTarget string `json:"promotion_target,omitempty"` // 业务字段	推广标的
		BiddingStrategy string `json:"bidding_strategy,omitempty"` // 业务字段	出价方式
		BuildType       string `json:"build_type,omitempty"`       // 业务字段	搭建类型
		MarketingTarget string `json:"marketing_target,omitempty"` // 业务字段	营销诉求
		PageId          string `json:"page_id,omitempty"`          // 业务字段	落地页id	带细分条件字段才有
		ItemId          string `json:"item_id,omitempty"`          // 业务字段	商品id	带细分条件字段才有
		LiveRedId       string `json:"live_red_id,omitempty"`      // 业务字段	直播间id	带细分条件字段才有
		CountryName     string `json:"country_name,omitempty"`     // 业务字段	国家	带细分条件字段才有
		Province        string `json:"province,omitempty"`         // 业务字段	省份	带细分条件字段才有
		City            string `json:"city,omitempty"`             // 业务字段	城市	带细分条件字段才有
	}

	// OfflineDataDTO 数据指标
	OfflineDataDTO struct {
		OfflineBizField
		DataReportDTO
	}

	// OfflineAccountDataDTO 表示账户层级离线报表数据
	OfflineAccountDataDTO struct {
		OfflineBizField
		DataReportDTO
	}

	// OfflineCampaignDataDTO 表示计划层级离线报表数据
	OfflineCampaignDataDTO struct {
		OfflineBizField

		DataReportDTO

		CampaignId   string `json:"campaign_id"`   // 业务字段	计划id
		CampaignName string `json:"campaign_name"` // 业务字段	计划名称
	}

	// OfflineKeywordDataDTO 表示关键词层级离线报表数据
	OfflineKeywordDataDTO struct {
		Keyword         string `json:"keyword"`               // 业务字段	关键词
		CampaignId      string `json:"campaign_id"`           // 业务字段	计划id
		CampaignName    string `json:"campaign_name"`         // 业务字段	计划名称
		UnitId          string `json:"unit_id"`               // 业务字段	单元id
		UnitName        string `json:"unit_name"`             // 业务字段	单元名称
		Time            string `json:"time"`                  // 业务字段	时间
		Placement       string `json:"placement"`             // 业务字段	广告类型
		OptimizeTarget  string `json:"optimize_target"`       // 业务字段	优化目标
		PromotionTarget string `json:"promotion_target"`      // 业务字段	推广标的
		BiddingStrategy string `json:"bidding_strategy"`      // 业务字段	出价方式
		BuildType       string `json:"build_type"`            // 业务字段	搭建类型
		MarketingTarget string `json:"marketing_target"`      // 业务字段	营销诉求
		PageId          string `json:"page_id,omitempty"`     // 业务字段	落地页id	带细分条件字段才有
		ItemId          string `json:"item_id,omitempty"`     // 业务字段	商品id	带细分条件字段才有
		LiveRedId       string `json:"live_red_id,omitempty"` // 业务字段	直播间id	带细分条件字段才有
		KeywordId       int64  `json:"keyword_id,omitempty"`  // 业务字段 关键词id
		DataReportDTO
	}

	// OfflineUnitDataDTO 表示单元层级离线报表数据
	OfflineUnitDataDTO struct {
		OfflineBizField

		DataReportDTO

		CampaignId   string `json:"campaign_id"`        // 业务字段	计划id
		CampaignName string `json:"campaign_name"`      // 业务字段	计划名称
		UnitId       string `json:"unit_id,omitempy"`   // 业务字段 单元id
		UnitName     string `json:"unit_name,omitempy"` // 业务字段 单元名称
	}

	// OfflineCreativeDTO 表示创意层级离线报表数据
	OfflineCreativeDTO struct {
		CampaignId      string `json:"campaign_id,omitempty"`      // 业务字段 计划id
		CampaignName    string `json:"campaign_name,omitempty"`    // 业务字段 计划名称
		UnitId          string `json:"unit_id,omitempty"`          // 业务字段 单元id
		UnitName        string `json:"unit_name,omitempty"`        // 业务字段 单元名称
		CreativityId    string `json:"creativity_id,omitempty"`    // 业务字段 创意id
		CreativityName  string `json:"creativity_name,omitempty"`  // 业务字段 创意名称
		CreativityImage string `json:"creativity_image,omitempty"` // 业务字段 创意图片
		NoteId          string `json:"note_id,omitempty"`          // 业务字段 笔记id
		OfflineBizField
		DataReportDTO
	}

	// OfflineSearchWordDTO 表示搜索词层级离线报表数据
	OfflineSearchWordDTO struct {
		SearchWord      string `json:"search_word,omitempy"`      // 业务字段 搜索词
		CampaignId      string `json:"campaign_id,omitempy"`      // 业务字段 计划id
		CampaignName    string `json:"campaign_name,omitempy"`    // 业务字段 计划名称
		UnitId          string `json:"unit_id,omitempy"`          // 业务字段 单元id
		UnitName        string `json:"unit_name,omitempy"`        // 业务字段 单元名称
		CreativityId    string `json:"creativity_id,omitempy"`    // 业务字段 创意id
		NoteId          string `json:"note_id,omitempy"`          // 业务字段 笔记id
		Time            string `json:"time,omitempy"`             // 业务字段 时间
		Placement       string `json:"placement,omitempy"`        // 业务字段 广告类型
		OptimizeTarget  string `json:"optimize_target,omitempy"`  // 业务字段 优化目标
		PromotionTarget string `json:"promotion_target,omitempy"` // 字段业务 推广标的
		BiddingStrategy string `json:"bidding_strategy,omitempy"` // 字段业务 出价方式
		BuildType       string `json:"build_type,omitempy"`       // 字段业务 搭建类型
		MarketingTarget string `json:"marketing_target,omitempy"` // 字段业务 营销诉求

		DataReportDTO
	}
)

// FilterClause 过滤条件
type FilterClause struct {
	Column   string   `json:"column"`   // 筛选列 可筛选字段名 creativityId：创意ID fee：消耗金额 impression：展示次数 click：点击量 interaction：笔记互动量
	Operator string   `json:"operator"` // 操作符  可选操作符 >：大于 <：小于 in：等于
	Values   []string `json:"values"`   // 值
}
