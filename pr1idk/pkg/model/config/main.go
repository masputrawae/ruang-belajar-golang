package config

import (
	"encoding/json"
	"log"
	"os"
	"pr1idk/pkg/model"
)

type Repository interface {
	Load() model.Config
}

type API struct {
	JsonFile string
	Config   model.Config
	Server   model.ServerConfig
	DB       model.DBConfig
}

func New(a *API) Repository {
	f, err := os.ReadFile(a.JsonFile)
	if err != nil {
		log.Fatal("\n/pkg/model/config/main.go\nError in line [20]:\n", err)
	}
	if err = json.Unmarshal(f, &a.Config); err != nil {
		log.Fatal("\n/pkg/model/config/main.go\nError in line [27]:\n", err)
	}
	return &API{JsonFile: a.JsonFile, Config: a.Config, Server: a.Server, DB: a.DB}
}

func (a *API) Load() model.Config {
	return a.Config
}
