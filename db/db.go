package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

var instance *DB = nil

func GetInstance() (*sql.Conn, error) {
	if instance == nil {
		instance.Load()
	}

	dur := time.Duration(instance.ConnTimeout) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), dur)
	return instance.db.Conn(ctx)
}

type DB struct {
	User        string
	Password    string
	Address     string
	ConnTimeout int

	db *sql.DB
}

func (db *DB) Load() error {
	if instance != nil {
		return errors.New("already database instance")
	}

	cfg := &mysql.Config{
		User:      db.User,
		Passwd:    db.Password,
		Net:       "tcp",
		Addr:      db.Address,
		DBName:    "boardsvr",
		ParseTime: true,
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
