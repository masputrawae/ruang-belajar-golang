package handler

import (
	"net/http"
	"pr2idk/pkg/repository"
	"pr2idk/pkg/utils"
	"pr2idk/view"

	"github.com/a-h/templ"
)

type Handler struct {
	Repo repository.RepoManage
}

func New(r repository.RepoManage) Handler {
	return Handler{Repo: r}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	view.App(templ.URL("/register"))

	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		hashed, err := utils.GenHashPassword(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		password = ""

		h.Repo.AddUser(r.Context(), username, hashed)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	view.App(templ.URL("/login"))

	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		user, err := h.Repo.FindUserByUsername(r.Context(), username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if utils.CheckHashPassword(user.Password, password) {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	}
}
