package main

import (
	"net/http"
	"pr1idk/pkg/config"
	"pr1idk/pkg/database"
	"pr1idk/pkg/handler"
	"pr1idk/pkg/repository"
)

func main() {
	cfg := config.New("config.json")
	db := database.New(cfg.DB)
	repo := repository.New(db)
	r := handler.New(handler.HandlerManage{Repo: repo})
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.GetTodos)
	http.ListenAndServe(":"+cfg.Server.Port, mux)
}
