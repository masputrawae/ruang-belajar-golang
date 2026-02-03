package handler

import (
	"net/http"
	"pr2todo/internal/model"
	"pr2todo/view/templates/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pages.Index(r.URL.Path, true, "", model.User{}).Render(r.Context(), w)
}
