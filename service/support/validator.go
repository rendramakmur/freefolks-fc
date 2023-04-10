package support

import (
	"errors"
	"unicode"
)

func ValidatePassword(password *string) ([]string, error) {
	var (
		messages     []string
		hasLowercase bool
		hasUppercase bool
		hasNumber    bool
	)

	if password == nil {
		return messages, errors.New("Password required")
	}

	for _, char := range *password {
		switch {
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasLowercase {
		messages = append(messages, "Password need to contains lower case")
	}

	if !hasUppercase {
		messages = append(messages, "Password need to contains upper case")
	}

	if !hasNumber {
		messages = append(messages, "Password need to contains number")
	}

	if len(*password) < 5 {
		messages = append(messages, "Password length must be more than 5 character")
	}

	if len(*password) >= 16 {
		messages = append(messages, "Password length must be lower than 16 character")
	}

	if hasLowercase && hasUppercase && hasNumber && len(*password) > 5 && len(*password) <= 16 {
		return messages, nil
	}

	return messages, errors.New("Password need a valid length & combination")
}
