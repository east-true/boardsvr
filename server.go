package main

import (
	"boardsvr/db"
	"boardsvr/handler"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var headers map[string]string = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

type Server struct {
	ListenPort string
	db         *db.DB
}

func (svr *Server) Run(db *db.DB) {
	svr.db = db

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(cors.Default())
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
	base := engine.Group("/api")
	handler.Board(base)

	engine.Run(":50007")
}
