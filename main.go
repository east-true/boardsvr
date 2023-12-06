package main

import (
	"boardsvr/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

var headers map[string]string = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

func main() {
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(func(ctx *gin.Context) {
		for key, val := range headers {
			reqVal := ctx.Request.Header.Get(key)
			if reqVal != val {
				ctx.Status(http.StatusBadRequest)
				return
			}
		}
	})

	engine.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"main": "hello"})
	})

	group := engine.Group("/api/board")
	handler.Board(group)

	engine.Run(":50007")
}
