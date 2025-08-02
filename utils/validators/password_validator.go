package validators

import (
	"fmt"
	"strings"
	"unicode"
)

const specialChars = "!@#$&"

func ValidatePassword(password string) (bool, error) {

	var valid bool = false
	var err error

	if len(password) < 8 {
		return false, PassordValidationError{
			Code:    "TOO_SHORT",
			Message: "Password must be at least 8 characters",
		}
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true

		case unicode.IsLower(c):
			hasLower = true

		case unicode.IsDigit(c):
			hasDigit = true

		case strings.ContainsRune(specialChars, c):
			hasSpecial = true
		}
	}
	valid = valid && hasDigit && hasLower && hasUpper && hasSpecial
	return valid, err

}

type PassordValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (m PassordValidationError) Error() string {
	return fmt.Sprintf("[%s] - %s", m.Code, m.Message)
}
