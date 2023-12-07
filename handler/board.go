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
		board.POST("", addBoard)
		board.PUT("/:board_id", editBoard)
		board.DELETE("/delete/:board_id", removeBoard)
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

func addBoard(ctx *gin.Context) {
	obj := new(dto.Board)
	if err := ctx.Bind(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		// TODO : insert

		ctx.JSON(http.StatusOK, nil)
	}
}

func editBoard(ctx *gin.Context) {
	obj := new(dto.Board)
	if err := ctx.Bind(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		// TODO : insert

		ctx.JSON(http.StatusOK, nil)
	}
}

func removeBoard(ctx *gin.Context) {
	val := ctx.Request.URL.Query().Get("board_id")
	if val == "" {
		ctx.Status(http.StatusUnprocessableEntity)
		return
	}

	// TODO : delete

	ctx.Status(http.StatusOK)
}
