package main

import (
	"boardsvr/db"
)

func main() {
	svr := new(Server)
	db, err := db.New(&db.Config{
		Driver:   db.DRIVER_MYSQL,
		Username: "root",
		Pwd:      "",
		IP:       "127.0.0.1",
		Port:     db.MYSQL_DEFAULT_PORT,
		DB:       "boardsvr",
	})
	if err != nil {
		panic(err)
	}

	svr.Run(db)
}
