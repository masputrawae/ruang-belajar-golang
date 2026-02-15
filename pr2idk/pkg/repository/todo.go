package repository

import "context"

func (r *Repo) AddTodo(ctx context.Context, task string) error {
	stmt := "INSERT INTO todos (task) VALUES (?)"
	_, err := r.DB.ExecContext(ctx, stmt, task)
	return err
}

func (r *Repo) UpdateTodoStatus(ctx context.Context, status bool, id int) error {
	stmt := "UPDATE todos SET done = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, status, id)
	return err
}

func (r *Repo) UpdateTodoTask(ctx context.Context, task string, id int) error {
	stmt := "UPDATE todos SET task = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, task, id)
	return err
}

func (r *Repo) RemoveTodo(ctx context.Context, id int) error {
	stmt := "DELETE todos FROM id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, id)
	return err
}
