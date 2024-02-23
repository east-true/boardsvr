package handler

import (
	"boardsvr/server/model"
	"context"
	"fmt"
	"net/http"

	"github.com/east-true/auth-go/jwt"
	"github.com/east-true/auth-go/jwt/claims"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (h *Handler) Login(ctx *gin.Context) {
	obj := new(model.UserDTO)
	ctx.BindJSON(obj)

	entity, err := h.user.SelectUserByID(obj.ID)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	dto := entity.ToDTO()
	if dto.Pwd != obj.Pwd {
		ctx.Status(http.StatusForbidden)
		return
	}

	token := jwt.NewAuthToken(dto.Role)
	access, _, err := token.GetTokens()
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Header("authorization", access)
	ctx.Status(http.StatusOK)
}

func (h *Handler) Logout(ctx *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer rdb.Close()
	val, ok := ctx.Get("claim")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	} else {
		claim, ok := val.(*claims.Claims)
		if !ok {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		result := rdb.Del(context.Background(), claim.Subject)
		if result.Err() != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) AddUser(ctx *gin.Context) {
	obj := new(model.UserDTO)
	ctx.BindJSON(obj)

	err := h.user.InsertUser(obj)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) EditUser(ctx *gin.Context) {
	obj := new(model.UserDTO)
	ctx.BindJSON(obj)

	err := h.user.UpdateUser(obj)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) RemoveUser(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindQuery(obj)

	err := h.user.DeleteUser(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
