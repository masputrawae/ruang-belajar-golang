package repository

import (
	"context"
	"database/sql"
	"pr3idk/internal/model"
)

type Repo struct {
	DB *sql.DB
}

type RepoManage interface {
	UserAdd(ctx context.Context, uN, uP string) error
	UserDelete(ctx context.Context, id int) error
	UserFindByID(ctx context.Context, id int) (model.User, error)
	UserFindByUsername(ctx context.Context, uN string) (model.User, error)

	TodoAdd(ctx context.Context, uID int, t, s, p string) error
	TodoAddTask(ctx context.Context, uID int, t string) error
	TodoUpdate(ctx context.Context, t, s, p string, id int) error
	TodoUpdateTask(ctx context.Context, t string, id int) error
	TodoUpdateStatus(ctx context.Context, s string, id int) error
	TodoUpdatePriority(ctx context.Context, p string, id int) error
	TodoDelete(ctx context.Context, id int) error
	TodoFindByUserID(ctx context.Context, id int) ([]model.Todo, error)
	MetaFindAll(ctx context.Context, metaType string) ([]model.Meta, error)
}

func New(db *sql.DB) RepoManage {
	return &Repo{DB: db}
}
