package utils

import "golang.org/x/crypto/bcrypt"

// generate password
func GeneratePasswordHash(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// check password
func CheckPasswordHash(h, p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p)) == nil
}
