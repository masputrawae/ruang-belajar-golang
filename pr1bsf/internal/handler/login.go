package handler

import (
	"fmt"
	"net/http"
	"pr1bsf/internal/model/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := user.Users{}

	uName := r.FormValue("username")
	uPass := r.FormValue("username")

	usr, err := data.GetUser(uName)
	if err != nil {
		fmt.Printf("uName: %v | Tidak Ditemukan", uName)
	}

	if ok := user.CheckPasswordHash(usr.Password, uPass); !ok {
		fmt.Println("Password Salah")
	}
}
