package handler

import (
	"boardsvr/server/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBoardList(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindQuery(obj)
	var entitys []*model.BoardEntity
	var err error
	if obj.Author == "" {
		entitys, err = h.board.SelectBoardAll()
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
	} else {
		entitys, err = h.board.SelectBoardByAuthor(obj.Author)
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

func (h *Handler) GetBoard(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindUri(obj)
	entity, err := h.board.SelectBoardByID(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, entity.ToDTO())
}

func (h *Handler) AddBoard(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	if err := ctx.Bind(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		err = h.board.InsertBoard(obj)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

func (h *Handler) EditBoard(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	if err := ctx.BindJSON(obj); err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
	} else {
		err = h.board.UpdateBoard(obj)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func (h *Handler) RemoveBoard(ctx *gin.Context) {
	obj := new(model.BoardDTO)
	ctx.BindUri(obj)
	err := h.board.DeleteBoard(obj.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
