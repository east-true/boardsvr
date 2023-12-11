package main

import (
	"boardsvr/db"
	"boardsvr/server"

	"github.com/alecthomas/kong"
)

func main() {
	svr := &server.Server{
		ConfigPath: "./config/config.json",
		ListenPort: ":50007",
		Prefix:     "/api",

		DB: &db.DB{
			User:        "test",
			Password:    "testtest",
			Destination: "127.0.0.1:3306",
		},
	}

	ctx := kong.Parse(svr)
	ctx.Command()
	svr.Run()
}
