package handler

import (
	"boardsvr/server/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoardList(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindQuery(obj)
	var entitys []*model.BoardEntity
	var err error
	if obj.Author == "" {
		entitys, err = model.SelectBoardAll()
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
	} else {
		entitys, err = model.SelectBoardByAuthor(obj.Author)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	dtos := make([]*model.BoardDTO, len(entitys))
	for i, entity := range entitys {
		dtos[i] = entity.ToDTO()
	}
	ctx.JSON(http.StatusOK, dtos)
}

func GetBoard(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindUri(obj)
	entity, err := model.SelectBoardByID(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, entity.ToDTO())
}

func AddBoard(ctx *gin.Context) {
	obj := new(model.BoardDTO)
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
	obj := new(model.BoardDTO)
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
	obj := new(model.BoardDTO)
	ctx.BindUri(obj)
	err := model.DeleteBoard(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
