package loggerx

import (
	"bytes"
	"context"
	"io"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 可以被引入的中间件

// 自动设置请求ID
func SetTradeId(ctx context.Context) context.Context {
	return context.WithValue(ctx, "trade_id", uuid.NewV4().String())
}

// 中间件 gin框架保存请求相关信息
func MiddlewareGetGinParams(log *Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buf := &bytes.Buffer{}
		tea := io.TeeReader(ctx.Request.Body, buf)
		body, err := io.ReadAll(tea)
		if err != nil {
			// panic(err)
		}

		// 截取body的前1000个字符
		bodys := string(body)
		if len(bodys) > 1000 {
            bodys = bodys[:1000]
        }

		ctx.Request.Body = io.NopCloser(buf)

		m := map[string]interface{}{
			"method": ctx.Request.Method,
			"uri":    ctx.Request.RequestURI,
			"body":   bodys,
			"query":  ctx.Request.URL.Query(),
			"header": ctx.Request.Header,
		}
		log.Infof(ctx, "request %+v", m)
		ctx.Next()

	}
}
