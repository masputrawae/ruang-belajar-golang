package repositories

import (
	"context"
	"database/sql"
	"pr1idk/internal/models"
)

type Repositories struct {
	DB *sql.DB
}

// user manage
type UserManage interface {
	UserCreate(ctx context.Context, u models.User) error
	UserFindByUsername(ctx context.Context, u string) (models.User, error)
}

func NewUser(db *sql.DB) UserManage {
	return &Repositories{DB: db}
}

// todo manage
type TodoManage interface {
	// TodoFindAll() ([]models.APITodo, error)
	// TodoFindByStatus() ([]models.APITodo, error)
	// TodoFindByPriority() ([]models.APITodo, error)
	// TodoCreate() error
	// TodoUpdate() error
	// TodoDelete() error
}

func NewTodo(db *sql.DB) TodoManage {
	return &Repositories{DB: db}
}
