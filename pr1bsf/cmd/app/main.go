package main

import (
	"fmt"
	"net/http"
	"pr1bsf/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", handler.HomeHandler)

	fmt.Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
