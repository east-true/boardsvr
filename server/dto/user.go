package dto

import (
	"database/sql"
)

type User struct {
	ID      string       `json:"user_id"`
	Pwd     string       `json:"user_pwd"`
	Email   string       `json:"user_email"`
	Created sql.NullTime `json:"user_created"`
}
