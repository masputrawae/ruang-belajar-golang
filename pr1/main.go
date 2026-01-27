package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tmpl struct {
	Username string
}

func HashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(h, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}

func CreateUser(u, e, p string) (bool, error) {
	data, err := LoadData("users.json")
	if err != nil {
		return false, err
	}

	for _, v := range data {
		if v.Username == strings.ToLower(u) || v.Email == e {
			return false, nil
		}
	}

	hPass, err := HashPassword(p)
	if err != nil {
		return false, nil
	}

	data = append(data, User{
		Username: strings.ToLower(u),
		Email:    e,
		Password: hPass,
	})

	dataJson, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return false, err
	}

	if err := os.WriteFile("users.json", dataJson, 0644); err != nil {
		return false, err
	}

	return true, nil
}

func LoadData(f string) ([]User, error) {
	var users []User
	file, err := os.ReadFile(f)
	if err != nil {
		return []User{}, err
	}

	if err := json.Unmarshal(file, &users); err != nil {
		return []User{}, err
	}

	return users, nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		password := r.FormValue("password")
		username := r.FormValue("username")

		users, err := LoadData("users.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var vPass string

		for _, v := range users {
			if v.Username == username {
				vPass = v.Password
			} else {
				err := errors.New("User Tidak Ditemukan")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		valid := CheckPassword(vPass, password)
		userPage := fmt.Sprintf("/user/%s", username)
		if valid {
			http.Redirect(w, r, userPage, http.StatusSeeOther)
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := map[string]string{
			"Title": "Home",
		}

		tmpl.Execute(w, data)
	})

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("user.html"))
		data := map[string]string{
			"Title": "Register",
		}
		tmpl.Execute(w, data)

		r.ParseForm()
		password := r.FormValue("password")
		email := r.FormValue("email")
		username := r.FormValue("username")

		valid, err := CreateUser(password, email, username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		userPage := fmt.Sprintf("/user/%s", username)

		if valid {
			http.Redirect(w, r, userPage, http.StatusSeeOther)
		}
	})

	r.Get("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		tmpl := template.Must(template.ParseFiles("user.html"))
		data := map[string]string{
			"Username": username,
		}

		tmpl.Execute(w, data)
	})

	log.Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
