package service

import (
	"context"
	"pr3idk/internal/model"
	"pr3idk/internal/repository"
)

type Service struct {
	Repo repository.RepoManage
}

type ServiceManage interface {
	UserAdd(ctx context.Context, uN, pN *string) error
	UserAuth(ctx context.Context, uN, pN string) (bool, error)

	TodoAdd(ctx context.Context, uID int, t, s, p *string) error
	TodoDelete(ctx context.Context, id int) error
	TodoFindByUserID(ctx context.Context, id int) (model.UserTodoAPI, error)
}

func New(r repository.RepoManage) ServiceManage {
	return &Service{Repo: r}
}
