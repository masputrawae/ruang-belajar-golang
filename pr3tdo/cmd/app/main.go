package main

import (
	"fmt"
	"net/http"
	"pr3tdo/view/pages"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir("./assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	// Routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home("Home", "Ini Deskripsi").Render(r.Context(), w)
	})

	router.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		pages.About("About", "Ini Deskripsi Di About").Render(r.Context(), w)
	})

	fmt.Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
