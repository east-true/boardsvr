package db_test

import (
	. "boardsvr/db"
	"context"
	"testing"

	"github.com/go-sql-driver/mysql"
)

// TODO : test
func TestDB(t *testing.T) {
	cfg := mysql.NewConfig()
	cfg.User = "test"
	cfg.Passwd = "testtest"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "boardsvr"

	if err := Load(cfg); err != nil {
		t.Error(err)
		return
	}

	if db, err := GetInstance(); err != nil {
		t.Error(err)
		return
	} else {
		ctx := context.Background()
		conn, err := db.Conn(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		conn.PingContext(ctx)

		conn.Close()
		db.Close()
	}
}
