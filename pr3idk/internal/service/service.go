package service

import (
	"context"
	"errors"
	"pr3idk/internal/model"
	"pr3idk/internal/repository"
	"pr3idk/internal/utils"
)

type Service struct {
	Repo repository.RepoManage
}

type ServiceManage interface {
	UserAdd(ctx context.Context, uN, pN *string) error
	UserAuth(ctx context.Context, uN, pN string) (bool, error)

	TodoAdd(ctx context.Context, uID int, t, s, p *string) error
	TodoFindByUserID(ctx context.Context, id int) (model.UserTodoAPI, error)
}

func New(r repository.RepoManage) ServiceManage {
	return &Service{Repo: r}
}

// SERVICE USER: add new user
func (s *Service) UserAdd(ctx context.Context, uN, pN *string) error {
	// check nil value
	if uN == nil && pN == nil {
		return errors.New("Password & Username tidak boleh kosong")
	}

	// check username
	username := utils.NormalizeUsername(*uN)

	// hash password
	hashed, err := utils.GenHashPassword(*pN)
	if err != nil {
		return err
	}

	if err := s.Repo.UserAdd(ctx, username, hashed); err != nil {
		return err
	}

	return nil
}

// SERVICE USER: auth
func (s *Service) UserAuth(ctx context.Context, uN, pN string) (bool, error) {
	username := utils.NormalizeUsername(uN)
	data, err := s.Repo.UserFindByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	if data.Username != uN && !utils.CheckHashPassword(data.Password, pN) {
		return false, errors.New("Username & Password tidak cocok")
	}

	return true, nil
}

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
