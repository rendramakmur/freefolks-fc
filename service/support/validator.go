package support

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"unicode"

	"github.com/rendramakmur/freefolks-fc/helper"
	"github.com/rendramakmur/freefolks-fc/repository"
)

func ValidatePassword(password *string) ([]string, error) {
	var (
		messages     []string
		hasLowercase bool
		hasUppercase bool
		hasNumber    bool
	)

	if password == nil {
		return messages, nil
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

func ValidateUniqueEmail(email *string, userRepository *repository.UserRepository) error {
	if email == nil {
		return nil
	}

	if isExist := userRepository.ExistByEmail(*email); isExist {
		return errors.New("Email is already used")
	}

	return nil
}

func ValidateAge(bd *time.Time) error {
	if bd == nil {
		return nil
	}

	if helper.GetAge(bd) > 65 {
		return errors.New("Age can not be more than 65")
	}

	if helper.GetAge(bd) < 10 {
		return errors.New("Age can not be less than 10")
	}

	return nil
}

func ValidateMobileNumber(number *string) error {
	if number == nil {
		return nil
	}

	pattern := `^0\d{9,12}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(*number) {
		return errors.New("Should start with 0 and the length between 9 - 12 digits")
	}

	return nil
}

func ValidateName(name *string) error {
	if name == nil {
		return nil
	}

	if len(*name) < 2 {
		return errors.New("Name should be more than 2 characters")
	}

	if len(*name) >= 50 {
		return errors.New("Name should be less than 50 characters")
	}

	for _, r := range *name {
		if !unicode.IsLetter(r) {
			return errors.New("Name should be only contains letter")
		}
	}

	return nil
}

func ValidateGlobalParam(slug *string, codeId *int, globalParamRepo *repository.GlobalParamRepository) error {
	if slug == nil || codeId == nil {
		return nil
	}

	_, err := globalParamRepo.GetDefaultDataBySlugAndCodeId(*slug, *codeId)
	if err != nil {
		return errors.New(fmt.Sprintf("Global param for %s and code %d not found", *slug, *codeId))
	}

	return nil
}
