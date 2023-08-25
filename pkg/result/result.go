package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	OK                    = OfResult(http.StatusOK, "ok")
	BAD_REQUEST           = OfResult(http.StatusBadRequest, "bad request")
	UNAUTHORIZED          = OfResult(http.StatusUnauthorized, "unauthorized")
	FORBIDDEN             = OfResult(http.StatusForbidden, "forbidden")
	NOT_FOUND             = OfResult(http.StatusNotFound, "not found")
	REQUEST_TIMEOUT       = OfResult(http.StatusRequestTimeout, "request timeout")
	METHOD_NOT_ALLOWED    = OfResult(http.StatusMethodNotAllowed, "method not allowed")
	INTERNAL_SERVER_ERROR = OfResult(http.StatusInternalServerError, "internal server error")
	FAILED                = OfResult(10001, "service error")
)

type ResultWrapper struct {
	status  int
	message string
	data    any
}

// OfResult new a ResultWrapper
func OfResult(status int, message string) *ResultWrapper {
	return &ResultWrapper{
		status:  status,
		message: message,
	}
}

func (rw *ResultWrapper) WithMessage(message string) *ResultWrapper {
	rw.message = message
	return rw
}

func (rw *ResultWrapper) WithData(data any) *ResultWrapper {
	rw.data = data
	return rw
}

func (rw *ResultWrapper) JSON(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  rw.status,
		"message": rw.message,
		"data":    rw.data,
	})
}
