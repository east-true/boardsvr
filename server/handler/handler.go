package handler

import "boardsvr/server/model"

type Handler struct {
	board model.BoardAdaptor
	user  model.UserAdaptor
}

func New(model *model.Model) *Handler {
	return &Handler{
		board: model,
		user:  model,
	}
}
