package model

type User struct {
	ID string `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	Username string `json:"username"`
	Password string `json:"password"`
}
