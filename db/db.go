package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

var database *sql.DB = nil

func Load(cfg *mysql.Config) error {
	if db, err := sql.Open("mysql", cfg.FormatDSN()); err != nil {
		return err
	} else {
		database = db
		fmt.Println("connect", cfg.FormatDSN())
		return database.Ping()
	}

}

func GetInstance() (*sql.Conn, error) {
	if database == nil {
		return nil, errors.New("not loaded database instance")
	}

	duration := 10 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), duration)

	return database.Conn(ctx)
}
