package dto

import (
	"database/sql"
)

type User struct {
	ID      string
	Pwd     string
	Email   string
	Created sql.NullTime
}
