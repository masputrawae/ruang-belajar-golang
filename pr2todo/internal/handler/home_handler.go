package handler

import (
	"net/http"
	"pr2todo/view/templates/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pages.Index().Render(r.Context(), w)
}
