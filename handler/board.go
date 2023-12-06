package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Board(g *gin.RouterGroup) {
	g.GET("", getBoard)
}

type BoardDTO struct {
	id      int
	Title   string
	Content string
	Writer  string
}

func getBoard(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &BoardDTO{
		id:      0,
		Title:   "GET",
		Content: "Board",
		Writer:  "me",
	})

}
