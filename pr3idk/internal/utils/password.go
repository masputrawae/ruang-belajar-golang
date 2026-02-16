package utils

import "golang.org/x/crypto/bcrypt"

func GenHashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(p),
		bcrypt.DefaultCost,
	)
	return string(hash), err
}

func CheckHashPassword(h, p string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(h),
		[]byte(p),
	) == nil
}
