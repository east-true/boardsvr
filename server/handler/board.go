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
	var res []*model.BoardEntity
	var err error
	if obj.Author == "" {
		res, err = model.SelectBoardAll()
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
	} else {
		res, err = model.SelectBoardByAuthor(obj.Author)
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

func GetBoard(ctx *gin.Context) {
	obj := new(model.BoardEntity)
	ctx.BindUri(obj)
	res, err := model.SelectBoardByID(int(obj.Id.Int32))
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func AddBoard(ctx *gin.Context) {
	obj := new(model.BoardEntity)
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
	obj := new(model.BoardEntity)
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
	obj := new(model.BoardEntity)
	ctx.BindUri(obj)
	err := model.DeleteBoard(int(obj.Id.Int32))
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
