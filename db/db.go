package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

var instance *Mysql = nil

func GetInstance() (*sql.Conn, error) {
	if instance == nil {
		instance.Load()
	}

	dur := time.Duration(instance.ConnTimeout) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), dur)
	return instance.db.Conn(ctx)
}

type Mysql struct {
	User        string `json:"username"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	ConnTimeout int    `json:"connection_timeout"`

	db *sql.DB
}

func (db *Mysql) Load() error {
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

func (db *Mysql) Close() {
	db.db.Close()
}
