package handlers

import (
	"net/http"
	"pr4idk/internal/utils"
	"pr4idk/view"
)

// -- file css
func CssFileHandler(w http.ResponseWriter, r *http.Request) {
	css, err := utils.LoadStaticAssets(view.CssFile, "assets/css/main.css")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/css")
	w.Write(css)
}

// -- file js
func JsFileHandler(w http.ResponseWriter, r *http.Request) {
	js, err := utils.LoadStaticAssets(view.JsFile, "assets/js/main.js")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(js)
}
