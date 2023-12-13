package handler

import (
	"boardsvr/server/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	obj := new(model.User)
	ctx.BindJSON(obj)

	err := model.InsertUser(obj)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func EditUser(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}

func RemoveUser(ctx *gin.Context) {

	ctx.Status(http.StatusOK)
}
