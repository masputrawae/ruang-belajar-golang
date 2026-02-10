package todo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"pr1idk/pkg/model"
)

type Repository interface {
	FindALL() ([]model.Todo, error)
	FindByID(id int) (model.Todo, error)
	ToJson() ([]byte, error)
	Create(task string) error
	Delete(id int) error
	Update(task string, done bool, id int) error
}

type API struct {
	Todos []model.Todo
	DB    *sql.DB
}

func New(a *API) Repository {
	return a
}

func (a *API) FindALL() ([]model.Todo, error) {
	rows, err := a.DB.Query("SELECT id, task, done FROM todos")
	if err != nil {
		log.Default().Println("\n/pkg/model/todo/main.go\nError in line [22]:\n", err)
		return nil, err
	}
	var t model.Todo
	for rows.Next() {
		rows.Scan(&t.ID, &t.Task, &t.Done)
		a.Todos = append(a.Todos, t)
	}
	defer rows.Close()
	return a.Todos, nil
}

func (a *API) FindByID(id int) (model.Todo, error) {
	var t model.Todo

	a.DB.
		QueryRow("SELECT id, task, done FROM todos WHERE id = ?", id).
		Scan(&t.ID, &t.Task, &t.Done)

	if err := sql.ErrNoRows; err != nil {
		return t, errors.New("ID not found")
	}

	return t, nil
}

func (a *API) ToJson() ([]byte, error) {
	j, err := json.MarshalIndent(a.Todos, "", "    ")
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (a *API) Create(task string) error {
	_, err := a.DB.Exec("INSERT INTO todos (task) VALUES (?)", task)
	if err != nil {
		return err
	}
	return nil
}

func (a *API) Delete(id int) error {
	_, err := a.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (a *API) Update(task string, done bool, id int) error {
	_, err := a.DB.Exec("UPDATE todos SET task = ?, done = ? WHERE id = ?", task, done, id)
	if err != nil {
		return err
	}
	return nil
}
