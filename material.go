package adxhsmarket

// MaterialService 表示聚光素材管理服务
type MaterialService service

type (
	// DirectLink 直达连接
	DirectLink struct {
		Id         int64  `json:"id"`          // 直达链接的id
		Url        string `json:"url"`         // url内容
		Type       int    `json:"type"`        // 类型，1-deeplink， 2-ulk
		Status     int    `json:"status"`      // 状态，1-审核中，2-审核通过，3-审核拒绝
		RemarkName string `json:"remark_name"` // 备注名称
	}
)
