package handler

import (
	"boardsvr/server/dto"
	"net/http"
	"strings"

	"github.com/east-true/jwt-go/claims"
	"github.com/gin-gonic/gin"
)

func TokenVerify(ctx *gin.Context) {
	val := ctx.Request.Header.Get("authorization")
	if strings.HasPrefix(val, "Bearer") {
		token := strings.Split(val, " ")[1]
		claim := new(claims.Claims)
		if claim.Verify(token) {
			return
		}
	}

	ctx.AbortWithStatus(http.StatusForbidden)
}

func Login(ctx *gin.Context) {
	obj := new(dto.User)
	ctx.BindJSON(obj)

	// TODO : valid pwd by id

	ctx.Status(http.StatusOK)
}

func Logout(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}
