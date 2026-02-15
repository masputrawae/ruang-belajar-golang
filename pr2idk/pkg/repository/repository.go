package repository

import (
	"context"
	"database/sql"
	"pr2idk/pkg/model"
)

type Repo struct {
	DB *sql.DB
}

type RepoManage interface {
	// -- User
	AddUser(ctx context.Context, u, p string) error
	FindUserByUsername(ctx context.Context, username string) (model.User, error)
	RemoveUserByUsername(ctx context.Context, username string) error

	// -- user -> todo
	FindAllTodoByUsername(ctx context.Context, username string) ([]model.Todo, error)

	// -- Todo
	AddTodo(ctx context.Context, task string) error
	UpdateTodoStatus(ctx context.Context, status bool, id int) error
	UpdateTodoTask(ctx context.Context, task string, id int) error
}

func New(db *sql.DB) RepoManage {
	return &Repo{DB: db}
}
