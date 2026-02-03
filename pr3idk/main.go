package main

import (
	"fmt"
	"net/http"

	"pr3idk/view/pages"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Username string
	Password string
}

var DataUsers = []User{
	{
		Username: "user",
		Password: "userpass",
	},
}

var ValidToken = "token-rahasia"

// ===== MAIN FUNCTION =====
func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", HomeHandler)

	router.Get("/login", LoginPageHandler)
	router.Get("/register", LoginPageHandler)

	router.Post("/login", LoginHandler)
	router.Post("/register", RegisterHandler)

	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/dashboard", DashboardHandler)
	})

	fmt.Println("server: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

// ===== AUTH MIDDLEWARE =====
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value != ValidToken {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ===== HANDLER =====
// ===== HANDLER GET =====
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	pages.Dashboard("Dashboard").Render(r.Context(), w)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	pages.Login().Render(r.Context(), w)
}

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	pages.Register().Render(r.Context(), w)
}

// ==== HANDLER POST =====
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	for i := range DataUsers {
		if DataUsers[i].Username == username && DataUsers[i].Password == password {
			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: ValidToken,
				Path:  "/",
			})

			w.Header().Set("Authorization", ValidToken)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	http.Error(w, "Username atau Password Salah", http.StatusUnauthorized)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	for _, user := range DataUsers {
		if user.Username == username {
			http.Error(w, "Username sudah digunakan", http.StatusUnauthorized)
			return
		}
	}

	DataUsers = append(DataUsers, User{
		Username: username,
		Password: password,
	})
}
