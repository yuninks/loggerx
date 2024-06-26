package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 设置普通的traceId
func SetTraceId(ctx context.Context, traceKey string) context.Context {
	if traceKey == "" {
		traceKey = "trace_id"
	}

	val := ctx.Value(traceKey)
	if val == nil {
		ctx = context.WithValue(ctx, traceKey, uuid.NewV4().String())
	}
	return ctx
}

// 设置Gin的traceId
func SetGinTraceId(traceKey string) gin.HandlerFunc {

	if traceKey == "" {
		traceKey = "trace_id"
	}

	return func(ctx *gin.Context) {
		traceId := ctx.Request.Header.Get(traceKey)
		if traceId == "" {
			traceId = uuid.NewV4().String()
		}
		ctx.Set(traceKey, traceId)
	}
}
