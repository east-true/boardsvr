package handler

import (
	"boardsvr/server/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO
func GetUser(ctx *gin.Context) {
	obj := new(dto.User)
	ctx.JSON(http.StatusOK, obj)
}

func AddUser(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}

func EditUser(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}

func RemoveUser(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}
