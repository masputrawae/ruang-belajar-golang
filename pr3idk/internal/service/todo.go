package service

import (
	"context"
	"pr3idk/internal/model"
)

// SERVICE TODO: create todo
func (sv *Service) TodoAdd(ctx context.Context, uID int, t, s, p *string) error {
	if s == nil && p == nil {
		if err := sv.Repo.TodoAddTask(ctx, uID, *t); err != nil {
			return err
		}
		return nil
	}
	sv.Repo.TodoAdd(ctx, uID, *t, *s, *p)
	return nil
}

// SERVICE TODO: delete todo
func (s *Service) TodoDelete(ctx context.Context, id int) error {
	return s.Repo.TodoDelete(ctx, id)
}

// SERVICE TODO: find todo
func (s *Service) TodoFindByUserID(ctx context.Context, id int) (model.UserTodoAPI, error) {
	var results model.UserTodoAPI

	user, err := s.Repo.UserFindByID(ctx, id)
	if err != nil {
		return results, err
	}

	todos, err := s.Repo.TodoFindByUserID(ctx, id)
	if err != nil {
		return results, err
	}

	results = model.UserTodoAPI{
		Username: user.Username,
		Todos:    todos,
	}

	return results, nil
}
