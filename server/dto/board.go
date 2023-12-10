package dto

import "time"

type Board struct {
	id      int
	Title   string
	Content string
	Author  User
	Created time.Time
}

func (b *Board) GetID() int {
	return b.id
}
