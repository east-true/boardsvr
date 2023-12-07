package db_test

import (
	. "boardsvr/db"
	"testing"
)

// TODO : test
func TestNew(t *testing.T) {
	cfg := &Config{
		Driver:   DRIVER_MYSQL,
		Username: "root",
		IP:       "127.0.0.1",
		Port:     MYSQL_DEFAULT_PORT,
		DB:       "boardsvr",
	}

	if db, err := New(cfg); err != nil {
		t.Error(err)
	} else {
		db.Close()
	}
}
