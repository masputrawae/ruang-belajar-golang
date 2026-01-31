package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Users []User

// Menambah User Baru
func (u *Users) NewUser(nU User) error {
	// Cek Apakah Username Sudah Ada?
	for _, user := range *u {
		if user.Username == nU.Username {
			return errors.New("Username sudah digunakan")
		}
	}

	// Normalisasi Username
	nUname := nU.Username
	if !strings.HasSuffix(nU.Username, "@") {
		nUname = fmt.Sprintf("%s%s", "@", nU.Username)
	}

	nID := uuid.New().String()
	nPass, err := GeneratePasswordHash(nU.Password)

	if err != nil {
		return err
	}

	(*u) = append((*u), User{
		ID:        nID,
		FirstName: nU.FirstName,
		LastName:  nU.LastName,
		Email:     nU.Email,
		Username:  nUname,
		Password:  nPass,
	})

	return nil
}

func (u Users) GetUser(uN string) (User, error) {
	var user User

	for _, usr := range u {
		if usr.Username == uN {
			user = usr
		} else {
			return user, errors.New("Username Tidak Ditemukan")
		}
	}

	return user, nil
}

// Membuat kata sandi yang di hash
func GeneratePasswordHash(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(p),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Mencocokkan hash dengan kata sandi
func CheckPasswordHash(h, p string) bool {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(h),
		[]byte(p),
	); err != nil {
		return false
	}

	return true
}
