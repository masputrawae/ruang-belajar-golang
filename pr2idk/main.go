package main

import (
	"log"
	"net/http"
	"pr2idk/pkg/config"
	"pr2idk/pkg/database"
	"pr2idk/pkg/handler"
	"pr2idk/pkg/repository"
	"pr2idk/view"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()
	cfg := config.New("config.toml")
	db := database.New(cfg.DB)
	repo := repository.New(db)
	h := handler.New(repo)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view.App(templ.SafeURL(r.URL.Path)).Render(r.Context(), w)
	})

	mux.HandleFunc("/register", h.Register)
	mux.HandleFunc("/login", h.Login)

	log.Println("Server running: http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
