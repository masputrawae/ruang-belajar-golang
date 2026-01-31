package user

import (
	"fmt"
	"testing"
)

var DummyUsers Users = Users{User{
	FirstName: "John",
	LastName:  "Doe",
	Email:     "johndoe@mail.com",
	Username:  "@johndoe",
	Password:  "JohndoePassword123",
}}

func TestNewUser(t *testing.T) {
	if err := DummyUsers.NewUser(User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@mail.com",
		Username:  "Johndoe1",
		Password:  "JohndoePassword123",
	}); err != nil {
		fmt.Println(err)
	}

	if err := DummyUsers.NewUser(User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@mail.com",
		Username:  "Johndoe",
		Password:  "JohndoePassword123",
	}); err != nil {
		fmt.Println(err)
	}
}

func TestGetUser(t *testing.T) {
	user, err := DummyUsers.GetUser("@johndoe")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
