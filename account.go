package adxhsmarket

import (
	"context"
	"net/http"
)

// AccountService 表示开放平台账户服务
type AccountService service

type (
	// AdCampaignTradeDetail 计划流水详情
	AdCampaignTradeDetail struct {
		CampaignId        int64  `json:"campaign_id"`         //	计划Id
		CampaignName      string `json:"campaign_name"`       //	计划名称
		LaunchDate        string `json:"launch_date"`         //	投放时间
		PayTime           string `json:"pay_time"`            //	操作时间
		CampaignDayBudget int64  `json:"campaign_day_budget"` //	日预算
		OrderAmount       int64  `json:"order_amount"`        //	消耗金额 单位: 分
	}

	// AccountBudgetData 账户日预算余额详情
	AccountBudgetData struct {
		TotalBalance            int64 `json:"total_balance"`             //	账户余额 单位: 分
		CashBalance             int64 `json:"cash_balance"`              //	现金余额 单位: 分
		ReturnBalance           int64 `json:"return_balance"`            //	常规返货余额 单位: 分
		CreditBalance           int64 `json:"credit_balance"`            //	授信金额 单位: 分
		FreezeBalance           int64 `json:"freeze_balance"`            //	冻结余额 单位: 分
		AvailableBalance        int64 `json:"available_balance"`         //	可用金额 单位: 分
		TodaySpend              int64 `json:"today_spend"`               //	今日花费 单位: 分
		CompensateReturnBalance int64 `json:"compensate_return_balance"` //	赔付返货余额 单位: 分
		AccountBudget           int64 `json:"account_budget"`            //	账户日预算 单位: 分
		LimitDayBudget          int64 `json:"limit_day_budget"`          //	是否限制预算	0-不限预算，1-限制预算
	}

	// AccountTradeDetail 账户流水详情
	AccountTradeDetail struct {
		LaunchDate       string `json:"launch_date"`        //	消费时间
		OperateType      int    `json:"operate_type"`       //	消费类型
		TradeTime        string `json:"trade_time"`         //	操作时间
		AccountName      string `json:"account_name"`       //	账户名称
		OrderAmount      int64  `json:"order_amount"`       //	消费金额
		Balance          int64  `json:"balance"`            //	余额
		TransferObject   string `json:"transfer_object"`    //	款项对象
		Remark           string `json:"remark"`             //	备注
		AccountType      int    `json:"account_type"`       //	资金类型
		AccountTypeName  string `json:"account_type_name"`  //	资金类型名称
		BusinessTypeName string `json:"business_type_name"` //	业务类型名称
	}

	// WhiteListData 账户白名单
	WhiteListData struct {
		InNoteForceBindSpuWhiteList bool `json:"in_note_force_bind_spu_white_list"` //  是否在笔记强绑spu白名单	true:在白名单中，此时选择得笔记需绑定spu，不然报错。
	}
)

type AccountAdOrderRequest struct {
	AdvertiserId int64  `json:"advertiser_id"` // 广告主Id
	PageNum      int    `json:"page"`          // 页数，默认1
	PageSize     int    `json:"page_size"`     // 页大小，默认20,最大 50
	StartTime    string `json:"start_time"`    // 开始时间，格式 yyyy-MM-dd	示例：2023-09-20
	EndTime      string `json:"end_time"`      // 结束时间，格式 yyyy-MM-dd	示例：2023-09-21
}

type AccountAdOrderData struct {
	Total int64                   `json:"total"`                    // 总记录数
	Spend int                     `json:"spend"`                    // 消费返货金额总计
	List  []AdCampaignTradeDetail `json:"ad_campaign_trade_detail"` // 计划流水详情
}

// ListAdOrders 获取账户计划流水
func (s *AccountService) ListAdOrders(ctx context.Context, req *AccountAdOrderRequest, options ...RequestOption) (*AccountAdOrderData, error) {
	path := "/api/open/jg/account/ad/order/info"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[AccountAdOrderData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AccountOrderRequest struct {
	AdvertiserId int64  `json:"advertiser_id"`           // 广告主Id
	PageNum      int    `json:"page,omitempty"`          // 页数，默认1
	PageSize     int    `json:"page_size,omitempty"`     // 页大小，默认20,最大100
	StartTime    string `json:"start_time"`              // 开始时间，格式 yyyy-MM-dd	示例：2023-09-20
	EndTime      string `json:"end_time"`                // 结束时间，格式 yyyy-MM-dd	示例：2023-09-21
	AccountTypes []int  `json:"account_types,omitempty"` // 资金账户类型		0-现金 1-常规返货 2-授信 6-赔付返货 不传，默认值为[0,1,2,6]
	Type         int    `json:"type,omitempty"`          // 数据类型	账户为服务商时必传，传type=2
}

type AccountOrderData struct {
	TotalCount     int64                `json:"total_count"`          // 总记录数
	AggregateOrder int                  `json:"aggregate_order"`      // 消费返货金额总计
	List           []AccountTradeDetail `json:"account_trade_detail"` // 计划流水详情
}

// ListOrders 获取账户流水
func (s *AccountService) ListOrders(ctx context.Context, req *AccountOrderRequest, options ...RequestOption) (*AccountOrderData, error) {
	path := "/api/open/jg/account/order/info"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[AccountOrderData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Budget 获取账户日预算余额
func (s *AccountService) Budget(ctx context.Context, advertiserId int32, options ...RequestOption) (*AccountBudgetData, error) {
	path := "/api/open/jg/account/budget/info"
	body := map[string]interface{}{"advertiser_id": advertiserId}

	response, err := s.client.Request(ctx, http.MethodPost, path, body, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[AccountBudgetData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// WhiteList 账户白名单
func (s *AccountService) WhiteList(ctx context.Context, advertiserId int32, options ...RequestOption) (*WhiteListData, error) {
	path := "/api/open/jg/white/list"
	body := map[string]interface{}{"advertiser_id": advertiserId}

	response, err := s.client.Request(ctx, http.MethodPost, path, body, nil, options...)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalApiResult[WhiteListData](response.RawBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}
