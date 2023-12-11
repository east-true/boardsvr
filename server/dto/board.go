package dto

import "database/sql"

type Board struct {
	Id      int          `json:"board_id" uri:"board_id"`
	Title   string       `json:"title"`
	Content string       `json:"content"`
	Author  string       `json:"author"`
	Updated sql.NullTime `json:"updated"`
}
