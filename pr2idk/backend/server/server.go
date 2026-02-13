package server

import (
	"pr2idk/backend/internal/configs"
	"pr2idk/backend/internal/databases"
)

func New() error {
	cfg := configs.Load("config.toml")
	db := databases.ConnectDB(cfg.Database)
	defer db.Close()

	return nil
}
