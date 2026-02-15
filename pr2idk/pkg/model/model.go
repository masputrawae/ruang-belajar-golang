package model

import "time"

type User struct {
	ID       int
	Username string
	Password string
}

type Todo struct {
	ID        int
	UserID    int
	Task      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
