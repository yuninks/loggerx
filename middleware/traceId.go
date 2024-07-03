package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/yuninks/loggerx"
)

// 设置普通的traceId
func SetTraceIdByKey(ctx context.Context, traceKey string) context.Context {
	if traceKey == "" {
		traceKey = "trace_id"
	}

	val := ctx.Value(traceKey)
	if val == nil {
		ctx = context.WithValue(ctx, traceKey, uuid.NewV4().String())
	}
	return ctx
}

// 设置logger的traceId
func SetTraceId(ctx context.Context, logger *loggerx.Logger) context.Context {
	return SetTraceIdByKey(ctx, logger.GetTraceField())
}

// 设置Gin的traceId
func SetGinTraceIdByKey(traceKey string) gin.HandlerFunc {

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

// 设置Gin的traceId
func SetGinTraceIdByLogger(logger *loggerx.Logger) gin.HandlerFunc {
	return SetGinTraceIdByKey(logger.GetTraceField())
}
