package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir("./assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	fmt.Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
