package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/chubin518/kestrel-layout-basic/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// RequestLoggingMiddleware Logs all requests to info log
func RequestLoggingMiddleware(skipRegexs ...string) gin.HandlerFunc {
	regexpPaths := make([]*regexp.Regexp, len(skipRegexs))
	for i, p := range skipRegexs {
		regexpPaths[i] = regexp.MustCompile(p)
	}
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}
		path := ctx.Request.URL.Path
		for _, regex := range regexpPaths {
			if regex.MatchString(path) {
				ctx.Next()
				return
			}
		}
		start := time.Now()

		var reqBody []byte
		if ctx.Request.Body != nil && ctx.Request.Body != http.NoBody {
			reqBody, _ = ctx.GetRawData()
			// 将原body塞回去
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}
		writer := &ResponseWriterWrapper{bodyBuf: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = writer

		ctx.Next()

		cost := time.Since(start)
		fields := []zapcore.Field{
			zap.Int64("time", cost.Milliseconds()),
			zap.Int("status", ctx.Writer.Status()),
			zap.String("reqbody", string(reqBody)),
			zap.String("respbody", writer.bodyBuf.String()),
		}
		if len(ctx.Errors) > 0 {
			fields = append(fields, zap.String("error", ctx.Errors.ByType(gin.ErrorTypePrivate).String()))
			logging.WithContext(ctx.Request.Context()).Info(fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.String()), fields...)
		} else {
			logging.WithContext(ctx.Request.Context()).Info(fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.String()), fields...)
		}
	}
}

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (rw *ResponseWriterWrapper) Write(buf []byte) (int, error) {
	if n, err := rw.bodyBuf.Write(buf); err != nil {
		return n, err
	}
	return rw.ResponseWriter.Write(buf)
}
