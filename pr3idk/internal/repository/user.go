package repository

import (
	"context"
	"pr3idk/internal/model"
)

func (r *Repo) UserAdd(ctx context.Context, uN, uP string) error {
	stmt := "INSERT INTO users(username, password) VALUES (?,?)"
	_, err := r.DB.ExecContext(ctx, stmt, uN, uP)
	return err
}

func (r *Repo) UserDelete(ctx context.Context, id int) error {
	stmt := "DELETE FORM users WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, id)
	return err
}

func (r *Repo) UserFindByID(ctx context.Context, id int) (model.User, error) {
	var u model.User
	stmt := "SELECT id, username, password FROM users WHERE id = ?"
	err := r.DB.QueryRowContext(ctx, stmt, id).Scan(&u.ID, &u.Username, &u.Password)
	return u, err
}

func (r *Repo) UserFindByUsername(ctx context.Context, uN string) (model.User, error) {
	var u model.User
	stmt := "SELECT id, username, password FROM users WHERE username = ?"
	err := r.DB.QueryRowContext(ctx, stmt, uN).Scan(&u.ID, &u.Username, &u.Password)
	return u, err
}
