package app

import (
	"fmt"
	"net/http"
	"pr2todo/internal/handler"
	"pr2todo/view/assets"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	favicon    = http.FS(assets.Favicon)
	css        = http.FS(assets.Css)
	javascript = http.FS(assets.Javascript)
)

// Server, membutuhkan (port, host)
func Server(port, host string) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Handle("/favicon.ico", http.FileServer(favicon))
	router.Handle("/main.css", http.FileServer(css))
	router.Handle("/main.js", http.FileServer(javascript))

	router.Get("/", handler.HomeHandler)

	fmt.Printf("\nServer running: http://%s%s\n\n", host, port)
	http.ListenAndServe(port, router)
}
