package utils

import "unicode"

func IsPasswordStrong(password string) bool {
	var hasNumber, hasUpper, hasLower, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasNumber && hasUpper && hasLower && hasSpecial && len(password) >= 8
}
