package utils

import (
	"fmt"
	"strings"
)

func NormalizeUsername(s string) string {
	uN := strings.ToLower(s)
	uN = strings.ReplaceAll(uN, " ", "_")
	if !strings.HasPrefix(uN, "@") {
		uN = fmt.Sprintf("%s%s", "@", uN)
	}
	return uN
}
