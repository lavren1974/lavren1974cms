package config

import (
	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	AppTitle string
	AppPort  string
}

func LoadConfig(path string) (*AppConfig, error) {
	var config AppConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
