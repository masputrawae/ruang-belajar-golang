package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(chi_middleware.Logger)

	// logger
	f, _ := os.OpenFile(".access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger := log.New(f, "", log.LstdFlags)

	router.Use(chi_middleware.RequestLogger(
		&chi_middleware.DefaultLogFormatter{
			Logger:  logger,
			NoColor: true,
		},
	))

	defer f.Close()

	// handler
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	fmt.Printf("===\nServer running: http://localhost:8080\nStart running: %s\n====\n", time.Now())
	http.ListenAndServe(":8080", router)
}
