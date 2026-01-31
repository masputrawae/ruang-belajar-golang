// internal/service/user.go
package service

import (
	"errors"
	"fmt"
	"pr1bsf/internal/config"
	"pr1bsf/internal/model/user"
	"time"
)

var users = LoadDataUser()

// Load Data Users
func LoadDataUser() user.Users {
	cfg := config.LoadConfig("config.yaml")
	users, err := user.LoadData(cfg.DataDir + "/" + cfg.UserData)
	if err != nil {
		fmt.Println(err)
	}

	return users
}

// Create New User
func CreateUser(newData user.User) error {
	// Check username
	if users.IsUsernameTaken(newData.Username) {
		return errors.New("Username already exists")
	}

	if !user.IsValidEmail(newData.Email) {
		return errors.New("Invalid email")
	}

	password, err := user.GenerateHashPassword(newData.Password)
	if err != nil {
		return err
	}

	users.Create(user.User{
		ID:       user.GenerateID(),
		Username: newData.Username,
		Password: password,

		FirstName: newData.FirstName,
		LastName:  newData.LastName,
		Email:     newData.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return nil
}

// Update / Edit User
func UpdateUser(newData user.User, username string) error {
	if users.IsUsernameTaken(newData.Username) {
		return errors.New("Username already exists")
	}

	if newData.Email != "" {
		if !user.IsValidEmail(newData.Email) {
			return errors.New("Invalid email")
		}
	}

	users.Update(newData, username)
	return nil
}
