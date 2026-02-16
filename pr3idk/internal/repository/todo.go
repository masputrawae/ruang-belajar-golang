package repository

import (
	"context"
	"fmt"
	"pr3idk/internal/model"
)

// TODO: add new (t: task, s: status id, p: priority id)
func (r *Repo) TodoAdd(ctx context.Context, uID int, t, s, p string) error {
	stmt := "INSERT INTO todos(user_id, task, status_id, priority_id) VALUES(?,?,?,?)"
	_, err := r.DB.ExecContext(ctx, stmt, uID, t, s, p)
	return err
}

// TODO: add new (t: task, uID: user id)
func (r *Repo) TodoAddTask(ctx context.Context, uID int, t string) error {
	stmt := "INSERT INTO todos(user_id, task) VALUES(?, ?)"
	_, err := r.DB.ExecContext(ctx, stmt, uID, t)
	return err
}

// TODO: update (t: task, s: status_id, p: priority_id, id: id todo)
func (r *Repo) TodoUpdate(ctx context.Context, t, s, p string, id int) error {
	stmt := "UPDATE todos SET task = ?, status_id = ?, priority_id = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, t, s, p, id)
	return err
}

// TODO: update task only (t: task, id: id todo)
func (r *Repo) TodoUpdateTask(ctx context.Context, t string, id int) error {
	stmt := "UPDATE todos SET task = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, t, id)
	return err
}

// TODO: update status todo (s: status_id, id: id todo)
func (r *Repo) TodoUpdateStatus(ctx context.Context, s string, id int) error {
	stmt := "UPDATE todos SET status_id = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, s, id)
	return err
}

// TODO: update priority todo (p: priority_id, id: todo id)
func (r *Repo) TodoUpdatePriority(ctx context.Context, p string, id int) error {
	stmt := "UPDATE todos SET priority_id = ? WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, p, id)
	return err
}

// TODO: delete todo (id: id todo)
func (r *Repo) TodoDelete(ctx context.Context, id int) error {
	stmt := "DELETE FROM todos WHERE id = ?"
	_, err := r.DB.ExecContext(ctx, stmt, id)
	return err
}

// TODO: find todos by user id (id: user id)
func (r *Repo) TodoFindByUserID(ctx context.Context, id int) ([]model.Todo, error) {
	var results []model.Todo

	stmt := `
		SELECT
			t.id, t.task, t.status_id, t.priority_id, t.created_at, t.updated_at
		FROM 
			todos AS t
		JOIN 
			users AS u
		ON t.user_id = ?
	`

	rows, err := r.DB.QueryContext(ctx, stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t model.Todo

		err := rows.Scan(&t.ID, &t.Task, &t.StatusID, &t.PriorityID, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, t)
	}

	return results, err
}

// Meta: (ctx: context, metaType: statuses or priorities)
func (r *Repo) MetaFindAll(ctx context.Context, metaType string) ([]model.Meta, error) {
	var results []model.Meta

	stmt := fmt.Sprintf("SELECT id, name FROM %s", metaType)
	rows, err := r.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s model.Meta

		err := rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return nil, err
		}

		results = append(results, s)
	}

	return results, err
}
