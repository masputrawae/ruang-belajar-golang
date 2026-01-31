package handler

import (
	"net/http"
	"pr1bsf/view/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
