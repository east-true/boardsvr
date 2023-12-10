package dto

import "time"

type User struct {
	ID      string
	Pwd     string
	Email   string
	Created time.Time
}
