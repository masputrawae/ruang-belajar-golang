package main

import (
	"fmt"
	"net/http"
	"pr4lgn/internal/data"
	"pr4lgn/internal/model"
	"pr4lgn/internal/model/user"
	"pr4lgn/view/pages"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var IsLogin = false
var Username = ""

func main() {
	userData, err := data.LoadUserData()
	if err != nil {
		fmt.Println("Data Tidak Ditemukan")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home(IsLogin, Username).Render(r.Context(), w)
	})

	router.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		pages.Login().Render(r.Context(), w)
	})

	router.Get("/registration", func(w http.ResponseWriter, r *http.Request) {
		pages.Registration().Render(r.Context(), w)
	})

	router.Post("/registration-form", func(w http.ResponseWriter, r *http.Request) {
		firstName := r.PostFormValue("first-name")
		lastName := r.PostFormValue("last-name")
		email := r.PostFormValue("email")
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		fmt.Printf("%s | %s | %s | %s | %s\n", firstName, lastName, email, username, password)

		if err := userData.Append(model.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Username:  username,
			Password:  password,
		}); err != nil {
			w.Write([]byte(err.Error()))
		}

		if err := data.SaveUserData(userData); err != nil {
			w.Write([]byte(err.Error()))
		}

		IsLogin = true
		Username = firstName + " " + lastName
	})

	router.Post("/login-form", func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		for _, v := range userData {
			if v.Username == username {
				Username = v.FirstName + " " + v.LastName
				ok, _ := user.CheckPasswordHash(v.Password, password)
				if ok {
					IsLogin = true
				}
			}
		}
	})

	fmt.Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
