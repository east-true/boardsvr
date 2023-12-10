package db

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

var instance *DB = nil

func GetInstance() (*sql.Conn, error) {
	if instance == nil {
		return nil, errors.New("not loaded database instance")
	}

	duration := 10 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), duration)

	return instance.db.Conn(ctx)
}

func Parse(sql string) string {
	sqlLines := strings.Split(sql, "\n")
	for i, line := range sqlLines {
		sqlLines[i] = strings.TrimSpace(line)
	}
	return strings.Join(sqlLines, " ")
}

type DB struct {
	User        string
	Password    string
	Destination string

	db *sql.DB
}

func (db *DB) Load() error {
	if instance != nil {
		return errors.New("already database instance")
	}

	cfg := &mysql.Config{
		User:   db.User,
		Passwd: db.Password,
		Net:    "tcp",
		Addr:   db.Destination,
		DBName: "boardsvr",
	}

	if sqlDB, err := sql.Open("mysql", cfg.FormatDSN()); err != nil {
		return err
	} else {
		db.db = sqlDB
		instance = db
		return db.db.Ping()
	}
}

func (db *DB) Close() {
	db.db.Close()
}
