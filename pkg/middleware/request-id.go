package middleware

import (
	"net/http"
	"strings"

	"github.com/chubin518/kestrel-layout-basic/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const X_REQUEST_ID = "X-Request-ID"

// RequestIdMiddleware Adds an indentifier to the response using the X-Request-ID header. Passes the X-Request-ID value back to the caller if it's sent in the request headers.
func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}
		traceId := ctx.GetHeader(X_REQUEST_ID)
		if traceId == "" {
			traceId = strings.ReplaceAll(uuid.New().String(), "-", "")
		}
		ctx.Request.Header.Add(X_REQUEST_ID, traceId)
		// Set the id to ensure that the requestid is in the response
		ctx.Header(X_REQUEST_ID, traceId)
		traceCtx := logging.NewContext(ctx.Request.Context(), map[string]any{
			"traceid": traceId,
		})
		ctx.Request = ctx.Request.WithContext(traceCtx)
		ctx.Next()
	}
}
