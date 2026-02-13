package routers

import (
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {

	})

	return r
}
