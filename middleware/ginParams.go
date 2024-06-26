package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuninks/loggerx"
)

// 可以被引入的中间件

// 自动设置请求ID

// 中间件 gin框架保存请求响应相关信息
func SetGinParams(log *loggerx.Logger) gin.HandlerFunc {
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

		// 下面是响应参数缓存
		blw := &bodyParamsWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = blw

		req := struct {
			Method   string              `json:"method"`
			Path     string              `json:"path"`
			ClientIP string              `json:"client_ip"`
			Body     string              `json:"body"`
			Header   map[string][]string `json:"header"`
		}{
			Method:   ctx.Request.Method,
			Path:     ctx.FullPath(),
			ClientIP: ctx.ClientIP(),
			Body:     bodyStr,
			Header:   ctx.Request.Header,
		}

		log.Info(ctx, "request", req)

		ctx.Next()

		// 判断返回的值是否json
		contentType := ctx.Writer.Header().Get("Content-Type")

		if !strings.Contains(contentType, "application/json") {
			return
		}

		respData := blw.body.String()

		resp := struct {
			HttpCode int    `json:"http_code"`
			Response string `json:"response"`
		}{
			HttpCode: ctx.Writer.Status(),
			Response: respData,
			// Header:   ctx.Writer.Header(),
		}
		log.Info(ctx, "response", resp)

	}
}

type bodyParamsWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 真正输出的时候，它会调用这个输出方法
func (w bodyParamsWriter) Write(b []byte) (int, error) {
	// b = []byte(`{"code":0}`)
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
