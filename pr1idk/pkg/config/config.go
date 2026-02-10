package config

import (
	"encoding/json"
	"os"
	"pr1idk/pkg/helper"
)

type Configs struct {
	DB     DB     `json:"db"`
	Server Server `json:"server"`
}

type DB struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func New(fn string) Configs {
	var configs Configs
	f, err := os.ReadFile(fn)
	if err != nil {
		helper.ErrFatal(err)
	}
	if err = json.Unmarshal(f, &configs); err != nil {
		helper.ErrFatal(err)
	}
	return configs
}
