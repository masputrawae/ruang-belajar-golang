package models

import "time"

type APITodo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Status    TodoMeta  `json:"status"`
	Priority  TodoMeta  `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todo struct {
	ID          int       `json:"id"`
	Task        string    `json:"task"`
	Status_id   string    `json:"status_id"`
	Priority_id string    `json:"priority_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoMeta struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
