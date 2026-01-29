package model

import "time"

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Todo struct {
	ID        string    `json:"id"`
	Task      string    `json:"task"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserManage interface {
	Append(User)
	Remove(User)
	Edit(User)
}

type TodoManage interface {
	Append(Todo)
	Remove(Todo)
	Edit(Todo)
}
