package server

import (
	"boardsvr/db"
	"boardsvr/server/handler"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var headers map[string]string = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

type Server struct {
	ConfigPath string
	ListenPort string
	Prefix     string
	DB         *db.DB
}

func (svr *Server) Run() {
	if err := svr.DB.Load(); err != nil {
		panic(err)
	} else {
		defer svr.DB.Close()
	}

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

	router := engine.Group(svr.Prefix)
	board := router.Group("/board")
	{
		board.GET("/list", handler.GetBoardList)
		board.GET("", handler.GetBoard)
		board.POST("/", handler.AddBoard)
		board.PUT("/:board_id", handler.EditBoard)
		board.DELETE("/:board_id", handler.RemoveBoard)
	}

	engine.Run(svr.ListenPort)
}
