package repository

import (
	"context"
	"database/sql"
	"pr1idk/pkg/model"
)

type TodoManage interface {
	FindAll(ctx context.Context) ([]model.Todo, error)
	FindByID(ctx context.Context, id int) (model.Todo, error)
	Create(ctx context.Context, task string) error
	Update(ctx context.Context, t model.Todo) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	DB *sql.DB
}

func New(db *sql.DB) TodoManage {
	return &Repository{DB: db}
}

func (r *Repository) FindAll(ctx context.Context) ([]model.Todo, error) {
	var results []model.Todo

	stmt := "SELECT id, task, done FROM todos"
	rows, err := r.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Task, &t.Done); err == sql.ErrNoRows {
			return nil, err
		}

		results = append(results, t)
	}
	return results, nil
}

func (r *Repository) FindByID(ctx context.Context, id int) (model.Todo, error) {
	var t model.Todo

	stmt := "SELECT id, task, done FROM todos WHERE id = ?"
	row := r.DB.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(&t.ID, &t.Task, &t.Done); err == sql.ErrNoRows {
		return t, err
	}
	return t, nil
}

func (r *Repository) Create(ctx context.Context, task string) error {
	stmt := "INSERT INTO todos (task) VALUES (?)"
	_, err := r.DB.ExecContext(ctx, stmt, task)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, t model.Todo) error {
	stmt := "UPDATE todos SET task = ?, done = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, t.Task, t.Done, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	stmt := "DELETE FROM todos WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
