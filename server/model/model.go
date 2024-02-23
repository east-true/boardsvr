package model

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type Model struct {
	mysql.Config

	instance *sql.DB
}

func (m *Model) Open() {
	db, err := sql.Open("mysql", m.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	m.instance = db
}

func (m *Model) Close() {
	m.instance.Close()
}
