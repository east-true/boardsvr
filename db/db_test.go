package db_test

import (
	. "boardsvr/db"
	"testing"

	"github.com/go-sql-driver/mysql"
)

var cfg *mysql.Config = &mysql.Config{
	User:   "test",
	Passwd: "testtest",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "boardsvr",
}

// TODO : test
func TestLoad(t *testing.T) {

	if err := Load(cfg); err != nil {
		t.Error(err)
		return
	}

	Close()
}

func TestGetInstance(t *testing.T) {
	if err := Load(cfg); err != nil {
		t.Error(err)
		return
	} else {
		defer Close()
	}

	if conn, err := GetInstance(); err != nil {
		t.Error(err)
		return
	} else {
		defer conn.Close()
	}
}
