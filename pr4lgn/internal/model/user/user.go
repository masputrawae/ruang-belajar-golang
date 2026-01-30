package user

import (
	"errors"
	"fmt"
	"pr4lgn/internal/model"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users []model.User

var (
	ErrUsernameExists     = errors.New("username already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserManage interface {
	Append(u *Users) error
}

func (u *Users) Append(nU model.User) error {
	// cek apakan username sudah ada
	for _, v := range *u {
		if v.Username == nU.Username {
			return ErrUsernameExists
		}
	}

	// cek dan validasi username
	var nUsername string
	if !strings.HasPrefix(nU.Username, "@") {
		nUsername = fmt.Sprintf("%s%s", "@", nU.Username)
	} else {
		nUsername = nU.Username
	}

	// id unik acak
	nID := uuid.New().String()

	// hashing password
	nPassword, err := PasswordHash(nU.Password)
	if err != nil {
		return err
	}

	(*u) = append((*u), model.User{
		ID: nID,

		FirstName: nU.FirstName,
		LastName:  nU.LastName,
		Email:     nU.Email,

		Username: nUsername,
		Password: nPassword,
	})

	// hapus password plaintext
	nU.Password = ""
	nPassword = ""

	return nil
}

// Function ====
func PasswordHash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPasswordHash(h, p string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p)); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
