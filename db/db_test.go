package db_test

import (
	. "boardsvr/db"
	"testing"
)

var database *Mysql = &Mysql{
	User:     "test",
	Password: "testtest",
	Address:  "127.0.0.1:3306",
}

func TestLoad(t *testing.T) {
	if err := database.Load(); err != nil {
		t.Error(err)
		return
	} else {
		defer database.Close()
	}
}

func TestGetInstance(t *testing.T) {
	if err := database.Load(); err != nil {
		t.Error(err)
		return
	} else {
		defer database.Close()
	}

	if conn, err := GetInstance(); err != nil {
		t.Error(err)
		return
	} else {
		defer conn.Close()
	}
}
