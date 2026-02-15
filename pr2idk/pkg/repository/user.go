package repository

import (
	"context"
	"pr2idk/pkg/model"
)

func (r *Repo) AddUser(ctx context.Context, u, p string) error {
	stmt := "INSERT INTO users ( username, password) VALUES( ?, ? )"
	_, err := r.DB.ExecContext(ctx, stmt, u, p)
	return err
}

func (r *Repo) RemoveUserByUsername(ctx context.Context, username string) error {
	stmt := "DELETE FROM users WHERE username = ?"
	_, err := r.DB.ExecContext(ctx, stmt, username)
	return err
}

func (r *Repo) FindUserByUsername(ctx context.Context, username string) (model.User, error) {
	var u model.User
	stmt := "SELECT id, username, password FROM users WHERE username = ?"
	row := r.DB.QueryRowContext(ctx, stmt, username)
	if err := row.Scan(u.ID, u.Username, u.Password); err != nil {
		return u, err
	}
	return u, nil
}

func (r *Repo) FindAllTodoByUsername(ctx context.Context, username string) ([]model.Todo, error) {
	var results []model.Todo
	stmt := `
	SELECT 
		t.id, t.task, t.done, t.created_at, t.updated_at
	FROM 
		todos AS t
	JOIN 
		users AS u 
	ON 
		u.id = t.user_id
	WHERE
		u.id = ?
	`

	rows, err := r.DB.QueryContext(ctx, stmt, username)
	if err != nil {
		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var t model.Todo
		err := rows.Scan(&t.ID, &t.Task, &t.Done, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return results, err
		}
		results = append(results, t)
	}

	return results, nil
}
