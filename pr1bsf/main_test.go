package main

import (
	"fmt"
	"net/mail"
	"testing"
)

func TestEmail(t *testing.T) {
	valid, err := mail.ParseAddress("bg@example.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Addres: %s\nName: %s\n", valid.Address, valid.Name)
}
