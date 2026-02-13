package repositories

import (
	"context"
	"pr1idk/internal/models"
)

func (r *Repositories) UserFindByUsername(ctx context.Context, u string) (models.User, error) {
	var user models.User
	stmt := "SELECT id, username, password, email FROM users WHERE username = ?"

	row := r.DB.QueryRowContext(ctx, stmt, u)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repositories) UserCreate(ctx context.Context, u models.User) error {
	stmt := `
	INSERT INTO 
		users (username, password, email)
	VALUES 
		(?, ?, ?)
	`
	_, err := r.DB.ExecContext(ctx, stmt, u.Username, u.Password, u.Email)
	if err != nil {
		return err
	}

	return nil
}
