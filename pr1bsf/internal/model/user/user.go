// internal/model/user/user.go
package user

import (
	"encoding/json"
	"errors"
	"net/mail"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
type UserManage interface {
	Create(User)
	Update(User)
	Get(string) (User, error)
	IsUsernameTaken(string) bool
	SaveData() error
}

func (users *Users) Create(newData User) {
	(*users) = append((*users), newData)
}

// Username Checking
func (users Users) IsUsernameTaken(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

// Update User
func (users *Users) Update(newData User, username string) {
	for i := range *users {
		if (*users)[i].Username == username {
			if newData.Username != "" {
				(*users)[i].Username = newData.Username
			}
			if newData.Password != "" {
				(*users)[i].Password = newData.Password
			}
			if newData.FirstName != "" {
				(*users)[i].FirstName = newData.FirstName
			}
			if newData.LastName != "" {
				(*users)[i].LastName = newData.LastName
			}
			if newData.Email != "" {
				(*users)[i].Email = newData.Email
			}

			(*users)[i].UpdatedAt = time.Now()
		}
	}
}

// Get Spesific User by Username
func (users Users) Get(username string) (User, error) {
	for i, user := range users {
		if user.Username == username {
			return users[i], nil
		}
	}
	return User{}, errors.New("Username not found")
}

// Save User Data to Json
func (users Users) SaveData(file string) error {
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(file, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

// Load User Data from Json
func LoadData(file string) (Users, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var data Users
	if err := json.Unmarshal(f, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// Email Checking
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Generate Hash Password for User
func GenerateHashPassword(p string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(p),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Checking Hash Password for User
func CheckingHashPassword(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}

// Generate Random UUID for User
func GenerateID() string {
	return uuid.New().String()
}
