package helper

import (
	"net/http"
	"strings"

	"github.com/east-true/jwt-go/claims"
	"github.com/gin-gonic/gin"
)

func TokenVerify(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("authorization")
	if strings.HasPrefix(auth, "Bearer") {
		token := strings.Split(auth, " ")[1]
		claim := new(claims.Claims)
		if claim.Verify(token) {
			return
		}
	}

	ctx.AbortWithStatus(http.StatusForbidden)
}
