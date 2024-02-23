package server

import (
	"boardsvr/server/handler"
	"net/http"

	"github.com/east-true/auth-go/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var headers map[string]string = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}

type Server struct {
	ConfigPath string
	ListenPort string
	Prefix     string
}

func (svr *Server) Run() {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.Use(jwt.JwtVerify)
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
	router.POST("/login", handler.Login)
	router.POST("/logout", handler.Logout)
	board := router.Group("/board")
	{
		board.GET("", handler.GetBoardList)
		board.GET("/:board_id", handler.GetBoard)
		board.POST("", handler.AddBoard)
		board.PUT("", handler.EditBoard)
		board.DELETE("/:board_id", handler.RemoveBoard)
	}
	user := router.Group("/user")
	{
		user.POST("", handler.AddUser)
		user.PUT("", handler.EditUser)
		user.DELETE("", handler.RemoveUser)
	}

	engine.Run(svr.ListenPort)
}
