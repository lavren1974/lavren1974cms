package config

import (
	"github.com/BurntSushi/toml"
)

type AppConfigGlobal struct {
	CmsName    string
	CmsVersion string
}

func LoadConfigGlobal() (*AppConfigGlobal, error) {
	var config AppConfigGlobal
	if _, err := toml.DecodeFile("../../config.toml", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
