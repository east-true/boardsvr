package main

import (
	"boardsvr/db"
	"boardsvr/server"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
)

func main() {
	var confPath string
	if str, err := os.Executable(); err == nil {
		fmt.Println(err)
	} else {
		confPath = filepath.Join(filepath.Dir(str), "config", "config.json")
	}

	svr := &server.Server{
		ConfigPath: confPath,
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
