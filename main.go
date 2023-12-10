package main

import (
	"boardsvr/db"
	"boardsvr/server"
)

func main() {
	svr := &server.Server{
		ListenPort: ":50007",
		Prefix:     "/api",

		DB: &db.DB{
			User:        "test",
			Password:    "testtest",
			Destination: "127.0.0.1:3306",
		},
	}

	svr.Run()
}
