package test

import (
	"pr3idk/internal/utils"
	"testing"
)

func TestUtilsPasswordSuccess(t *testing.T) {
	p := "Ini Password"
	hash, err := utils.GenHashPassword(p)

	if err != nil {
		t.Error(err)
	}

	if ok := utils.CheckHashPassword(hash, p); ok {
		t.Error("Password Tidak Sama")
	}
}
