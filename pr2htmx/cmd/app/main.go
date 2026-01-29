package main

import (
	"net/http"
	"pr2htmx/internal/template"
)

var count = 0

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/counter", counterHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template.Layout(template.Index()).Render(r.Context(), w)
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		count++
	}
	template.Counter(count).Render(r.Context(), w)
}
