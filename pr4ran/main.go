package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

func main() {
	Users = append(Users, User{Name: "Jono", Age: 20})

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Welcome")
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&Users)
	})

	mux.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, _ := strconv.Atoi(id)

		if idInt >= len(Users)+1 || idInt == 0 {
			return
		}

		user := Users[idInt-1]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&user)
	})

	mux.HandleFunc("/user/add", func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			return
		}

		for i := range Users {
			if Users[i].Name == user.Name {
				return
			}
		}

		Users = append(Users, user)
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
