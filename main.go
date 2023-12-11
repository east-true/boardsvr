package main

import (
	"boardsvr/server"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var confPath string
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	dir := filepath.Dir(path)
	confPath = filepath.Join(dir, "config", "config.json")

	b, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	svr := new(server.Server)
	err = json.Unmarshal(b, svr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// svr := &server.Server{
	// 	ConfigPath: confPath,
	// 	ListenPort: ":50007",
	// 	Prefix:     "/api",

	// 	DB: &db.DB{
	// 		User:        "test",
	// 		Password:    "testtest",
	// 		Destination: "127.0.0.1:3306",
	// 	},
	// }

	svr.Run()
}
