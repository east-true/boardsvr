package handler

import (
	"boardsvr/server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	obj := new(model.User)
	ctx.BindJSON(obj)

	// TODO : valid pwd by id

	ctx.Status(http.StatusOK)
}

func Logout(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}
