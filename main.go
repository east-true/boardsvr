package main

import (
	"boardsvr/handler"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	svr := &Server{
		ListenPort: ":50007",
		Prefix:     "/api",
	}

	svr.Run()
}

var headers map[string]string = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

type Server struct {
	ListenPort string
	Prefix     string
}

func (svr *Server) Run() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
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
	base := engine.Group(svr.Prefix)
	handler.Board(base)

	engine.Run(svr.ListenPort)
}
