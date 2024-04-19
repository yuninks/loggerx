package loggerx

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// 可以被引入的中间件

// 自动设置请求ID
func SetTradeId(ctx context.Context) context.Context {
	return context.WithValue(ctx, "trade_id", uuid.NewV4().String())
}
