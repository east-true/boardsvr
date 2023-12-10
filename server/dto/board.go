package dto

import "time"

type Board struct {
	Id      int       `json:"board_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Created time.Time `json:"created"`
}
