package utils

import (
	"regexp"
)

func EmailValidate(email string) bool {
	var regex = `^[\w\.-]+@[\w\.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
