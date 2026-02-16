package utils

import (
	"fmt"
	"strings"
)

func UsernamePrefix(s string) string {
	username := s
	if strings.HasPrefix(s, "@") {
		username = fmt.Sprintf("%s%s", "@", s)
	}
	return username
}
