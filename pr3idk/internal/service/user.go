package service

import (
	"context"
	"errors"
	"pr3idk/internal/utils"
)

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
