package utils

func StrToPtr(s string) *string {
	return &s
}

func PtrToStr(s *string) string {
	return *s
}
