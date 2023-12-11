package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}

func Logout(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}
