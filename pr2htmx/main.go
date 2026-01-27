package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type Todos []Todo

var data Todos

func LoadData(fN string) error {
	file, err := os.ReadFile(fN)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return err
	}

	return nil
}

func SaveData(fN string, nD Todos) error {
	jData, err := json.MarshalIndent(nD, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(fN, jData, 0644); err != nil {
		return err
	}

	return nil
}

func (t *Todos) Sorted() {
	sort.Slice((*t), func(i, j int) bool { return (*t)[i].ID < (*t)[j].ID })
}

// Menambah Tugas Baru
func (t *Todos) Append(nT string) error {
	(*t).Sorted()
	nID := 1
	for _, v := range *t {
		if v.ID == nID {
			nID++
		}
	}

	(*t) = append((*t), Todo{
		ID:   nID,
		Task: nT,
	})

	if err := SaveData("data.json", *t); err != nil {
		return err
	}

	return nil
}

func (t *Todos) Remove(id int) error {
	for i, v := range *t {
		if v.ID == id {
			(*t) = append((*t)[:i], (*t)[i+1:]...)

			if err := SaveData("data.json", *t); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	if err := LoadData("data.json"); err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("view/index.html"))
		tmpl.Execute(w, "")
	})

	r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		newTask := r.FormValue("task")

		data.Append(newTask)
		data.Sorted()

		for _, v := range data {
			w.Write([]byte("<li>" + "ID: " + strconv.Itoa(v.ID) + " " + v.Task + "</li>"))
		}
	})

	r.Post("/load", func(w http.ResponseWriter, r *http.Request) {
		if err := LoadData("data.json"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data.Sorted()
		for _, v := range data {
			w.Write([]byte("<li>" + "ID: " + strconv.Itoa(v.ID) + " " + v.Task + "</li>"))
		}
	})

	r.Post("/remove", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		ids := r.FormValue("id")
		idint, _ := strconv.Atoi(ids)

		if err := data.Remove(idint); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data.Sorted()
		for _, v := range data {
			w.Write([]byte("<li>" + "ID: " + strconv.Itoa(v.ID) + " " + v.Task + "</li>"))
		}
	})

	fmt.Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
