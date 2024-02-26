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
	confPath = filepath.Join(dir, "arch", "config.json")

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

	svr.Run()
}
