package services

import (
	"context"
	"errors"
	"pr1idk/internal/models"
	"pr1idk/internal/repositories"
	"pr1idk/internal/utils"
)

type UserService struct {
	User repositories.UserManage
}

func NewUserService(r repositories.UserManage) UserService {
	return UserService{User: r}
}

func (uS *UserService) UserAuth(ctx context.Context, uName, uPass string) (bool, error) {
	user, err := uS.User.UserFindByUsername(ctx, uName)
	if err != nil {
		return false, err
	}
	if user.Username == uName && utils.CheckPasswordHash(user.Password, uPass) {
		return true, nil
	}

	return false, nil
}

func (uS *UserService) UserCreate(ctx context.Context, uName, uPass, uEmail string) error {
	hash, err := utils.GeneratePasswordHash(uPass)
	if err != nil {
		return err
	}

	if !utils.EmailValidate(uEmail) {
		return errors.New("invalid email")
	}

	if err := uS.User.UserCreate(ctx, models.User{
		Username: uName,
		Password: hash,
		Email:    uEmail,
	}); err != nil {
		return err
	}

	return nil
}
