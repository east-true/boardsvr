package handler

import (
	"boardsvr/server/helper/token"
	"boardsvr/server/model"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Login(ctx *gin.Context) {
	obj := new(model.UserDTO)
	ctx.BindJSON(obj)

	entity, err := model.SelectUserByID(obj.ID)
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

	token := token.NewAuthToken(dto.Role)
	access, _, err := token.GetTokens()
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Header("authorization", access)
	ctx.Status(http.StatusOK)
}

func Logout(ctx *gin.Context) {
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
		claim, ok := val.(*token.Claims)
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

func AddUser(ctx *gin.Context) {
	obj := new(model.UserDTO)
	ctx.BindJSON(obj)

	err := model.InsertUser(obj)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func EditUser(ctx *gin.Context) {
	obj := new(model.UserDTO)
	ctx.BindJSON(obj)

	err := model.UpdateUser(obj)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func RemoveUser(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindQuery(obj)

	err := model.DeleteUser(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
