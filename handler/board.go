package handler

import (
	"boardsvr/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Board(base *gin.RouterGroup) {
	board := base.Group("/board")
	{
		board.GET("", getBoard)
	}
}

func getBoard(ctx *gin.Context) {
	obj := &dto.Board{
		Title:   "GET",
		Content: "Board",
		Writer:  "me",
	}

	ctx.JSON(http.StatusOK, obj)
}
