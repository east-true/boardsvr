package main

import (
	"boardsvr/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"main": "hello"})
	})

	group := engine.Group("/api/board")
	handler.Board(group)

	engine.Run(":50007")
}
