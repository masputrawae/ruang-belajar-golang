package handler

import (
	"net/http"
	"pr2todo/internal/data"
	"pr2todo/internal/model"
	"pr2todo/internal/utils"
	"pr2todo/view/templates/pages"

	"github.com/google/uuid"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	message := ""

	if r.Method == http.MethodPost {
		username := r.FormValue("username")

		for i := range data.DataUsers {
			if data.DataUsers[i].Username == username {
				if ok := utils.CheckHashPassword(data.DataUsers[i].Password, r.FormValue("password")); ok != false {
					message = "Password Salah"
					pages.Index(r.URL.Path, false, message, model.User{}).Render(r.Context(), w)
				} else {
					message = "Success Login"
					pages.Index(r.URL.Path, false, message, data.DataUsers[i]).Render(r.Context(), w)
				}
			}
		}
	}
	pages.Index(r.URL.Path, false, message, model.User{}).Render(r.Context(), w)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	message := ""

	if r.Method == http.MethodPost {
		hashed, err := utils.GenerateHashPassword(r.FormValue("password"))

		if err != nil {
			message = err.Error()
		}

		id := uuid.New().String()
		firstName := r.FormValue("first-name")
		lastName := r.FormValue("last-name")
		username := r.FormValue("username")

		for i := range data.DataUsers {
			if data.DataUsers[i].Username != username {
				data.DataUsers = append(data.DataUsers, model.User{
					ID:        id,
					FirstName: firstName,
					LastName:  lastName,
					Username:  username,
					Password:  hashed,
				})

				data.DataSessions = append(data.DataSessions, model.Session{
					ID:       uuid.New().String(),
					Username: username,
				})

				pages.Index(r.URL.Path, true, message, data.DataUsers[i]).Render(r.Context(), w)
			} else {
				message = "Username sudah digunakan"
				pages.Index(r.URL.Path, false, message, model.User{}).Render(r.Context(), w)
			}
		}
	}
	pages.Index(r.URL.Path, false, message, model.User{}).Render(r.Context(), w)
}
