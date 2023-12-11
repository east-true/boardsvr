package config

import (
	"fmt"
	"os"
)

type Config struct {
	Path string `flag:"path"`
}

func (cfg *Config) Read() error {
	b, err := os.ReadFile(cfg.Path)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
