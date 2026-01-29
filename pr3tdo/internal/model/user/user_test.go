package user

import (
	"pr3tdo/internal/model"
	"testing"
)

func TestUsers_Append_Success(t *testing.T) {
	users := Users{}

	err := users.Append(model.User{
		FirstName: "Putra",
		LastName:  "Test",
		Username:  "putra",
		Password:  "secret",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}

	if users[0].Password == "secret" {
		t.Fatal("password should be hashed, got plaintext")
	}
}

func TestUsers_Append_DuplicateUsername(t *testing.T) {
	users := Users{}

	_ = users.Append(model.User{
		Username: "putra",
		Password: "secret",
	})

	err := users.Append(model.User{
		Username: "putra",
		Password: "another",
	})

	if err == nil {
		t.Fatal("expected error for duplicate username, got nil")
	}
}

func TestUsers_Authenticate_Success(t *testing.T) {
	users := Users{}

	_ = users.Append(model.User{
		Username: "putra",
		Password: "secret",
	})

	user, err := users.Authenticate("putra", "secret")
	if err != nil {
		t.Fatalf("expected success, got error %v", err)
	}

	if user.Username != "putra" {
		t.Fatalf("expected username putra, got %s", user.Username)
	}
}
