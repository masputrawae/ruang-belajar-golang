package user

import (
	"errors"
	"pr3tdo/internal/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUsernameExists     = errors.New("username already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Users []model.User

func (u *Users) Append(nU model.User) error {
	for _, user := range *u {
		if user.Username == nU.Username {
			return ErrUsernameExists
		}
	}

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(nU.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	id := uuid.New()

	*u = append(*u, model.User{
		ID:        id.String(),
		FirstName: nU.FirstName,
		LastName:  nU.LastName,
		Username:  nU.Username,
		Password:  string(hashed),
	})

	return nil
}

func (u Users) GetUser(username string) (model.User, error) {
	for _, user := range u {
		if user.Username == username {
			return user, nil
		}
	}
	return model.User{}, ErrInvalidCredentials
}

func (u Users) Authenticate(username, password string) (model.User, error) {
	for _, user := range u {
		if user.Username == username {
			if bcrypt.CompareHashAndPassword(
				[]byte(user.Password),
				[]byte(password),
			) == nil {
				return user, nil
			}
			break
		}
	}
	return model.User{}, ErrInvalidCredentials
}
