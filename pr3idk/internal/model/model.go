package model

import "time"

type UserTodoAPI struct {
	Username string
	Todos    []Todo
}

type User struct {
	ID       int
	Username string
	Password string
}

type Todo struct {
	ID         int
	Task       string
	StatusID   string
	PriorityID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Meta struct {
	ID   string
	Name string
}
