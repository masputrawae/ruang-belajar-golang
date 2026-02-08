package main

import (
	"fmt"
	"net/http"
	"pr1mdl/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome")
	})

	http.ListenAndServe(":8080", middleware.Logger(mux))
}
