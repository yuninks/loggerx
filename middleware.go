package loggerx

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 可以被引入的中间件

// 自动设置请求ID
func SetTraceId(ctx context.Context) context.Context {
	return context.WithValue(ctx, "trace_id", uuid.NewV4().String())
}

// 中间件 gin框架保存请求相关信息
func MiddlewareGetGinParams(log *Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 使用bytes.Buffer来读取并记录请求体，同时避免改变ctx.Request.Body
		buf := &bytes.Buffer{}
		tea := io.TeeReader(ctx.Request.Body, buf)

		// 读取body
		body, err := io.ReadAll(tea)
		if err != nil {
			log.Infof(ctx, "Error reading request body: %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// 截取body的前1000个字符
		bodyStr := string(body)
		if len(bodyStr) > 1000 {
			bodyStr = bodyStr[:1000]
		}

		// 使用NopCloser包裹buffer，仅为了确保在读取body之后body可以被关闭，但并不改变原始的Request.Body
		ctx.Request.Body = io.NopCloser(buf)

		m := map[string]interface{}{
			"method": ctx.Request.Method,
			"uri":    ctx.Request.RequestURI,
			"body":   bodyStr,
			"query":  ctx.Request.URL.Query(),
			"header": ctx.Request.Header,
		}
		log.Infof(ctx, "request %+v", m)
		ctx.Next()

	}
}
