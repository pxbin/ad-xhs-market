# 小红书开放平台

### 使用
```go
import "github.com/pxbin/ad-xhs-market"

client := adxhsmarket.NewClient(appId, secret)
// or
client := adxhsmarket.NewClient(appId, secret, options...)
```

### 获取Token
```go
token, err := client.Auth.AccessToken(ctx, authCode)
```

### 刷新Token
```go
token, err := client.Auth.RefreshToken(ctx, refreshToken)
```

### 账户服务

获取计划流水
```go
req := &AccountAdOrderRequest{
	AdvertiserId: int64(advertiserId),
	PageNum:      1,
	PageSize:     10,
	StartTime:    "2024-08-14",
	EndTime:      "2024-08-14",
}

result, err := adClient.Account.ListAdOrders(context.Background(), req, WithAccessToken(token))
```

获取账户日预算余额
```go
req := &AccountOrderRequest{
	AdvertiserId: int64(advertiserId),
	PageNum:      1,
	PageSize:     10,
	StartTime:    "2024-08-14",
	EndTime:      "2024-08-14",
	Type:         2,
}

result, err := adClient.Account.ListOrders(context.Background(), req, WithAccessToken(token))
```

账户白名单
```go
result, err := adClient.Account.WhiteList(context.Background(), int32(advertiserId), WithAccessToken(token))
```

### 数据报表

创意层级离线报表数据
```go
req := &OfflineCreativtyRequest{
	ListOptions: ListOptions{
		AdvertiserId: int64(advertiserId),
		PageNum:      1,
		PageSize:     10,
		StartDate:    "2024-08-13",
		EndDate:      "2024-08-13",
	},
}

result, err := client.Report.ListOfflineCreativty(context.Background(), req, WithAccessToken(token))
```