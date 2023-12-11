package handler

import (
	"boardsvr/server/dto"
	"boardsvr/server/model"
	"fmt"
	"net/http"

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
	obj := new(dto.Board)
	ctx.BindUri(obj)
	res, err := model.SelectBoardByID(obj.Id)
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
	if err := ctx.BindJSON(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		err = model.UpdateBoard(obj)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func RemoveBoard(ctx *gin.Context) {
	obj := new(dto.Board)
	ctx.BindUri(obj)
	err := model.DeleteBoard(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
