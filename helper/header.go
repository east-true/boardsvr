package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var headers map[string]string = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

func Header(ctx *gin.Context) {
	for key, val := range headers {
		reqVal := ctx.Request.Header.Get(key)
		if reqVal != val {
			ctx.Status(http.StatusBadRequest)
			return
		}
	}
}
