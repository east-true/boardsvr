package handler

import (
	"boardsvr/server/dto"
	"boardsvr/server/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBoardList(ctx *gin.Context) {
	res, err := model.SelectBoardAll()
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func GetBoard(ctx *gin.Context) {
	val := ctx.Request.URL.Query().Get("board_id")
	if val == "" {
		ctx.Status(http.StatusUnprocessableEntity)
		return
	}

	res, err := model.SelectBoardByID(val)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func AddBoard(ctx *gin.Context) {
	obj := new(dto.Board)
	if err := ctx.Bind(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		err = model.InsertBoard(obj)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

func EditBoard(ctx *gin.Context) {
	obj := new(dto.Board)
	if err := ctx.Bind(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		err = model.UpdateBoard(obj)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

func RemoveBoard(ctx *gin.Context) {
	val := ctx.Request.URL.Query().Get("board_id")
	if val == "" {
		ctx.Status(http.StatusUnprocessableEntity)
		return
	}

	id, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = model.DeleteBoard(id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
