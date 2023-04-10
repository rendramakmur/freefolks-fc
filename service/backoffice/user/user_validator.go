package user

import (
	"strings"

	"github.com/go-playground/validator/v10"
	backofficeRequest "github.com/rendramakmur/freefolks-fc/model/request/backoffice"
	"github.com/rendramakmur/freefolks-fc/repository"
	customValidator "github.com/rendramakmur/freefolks-fc/service/support"
)

func ValidateCreateUserRequest(request *backofficeRequest.CreateUserRequest, userRepository repository.UserRepository) ([]string, error) {
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		return []string{strings.TrimSpace(strings.Split(err.Error(), ":")[2])}, err
	}

	if msg, err := customValidator.ValidatePassword(request.Password); err != nil {
		return msg, err
	}

	if err := customValidator.ValidateUniqueEmail(request.Email, &userRepository); err != nil {
		return []string{err.Error()}, err
	}

	if err := customValidator.ValidateAge(request.DateOfBirth.Time); err != nil {
		return []string{err.Error()}, err
	}

	if err := customValidator.ValidateMobileNumber(request.MobileNumber); err != nil {
		return []string{err.Error()}, err
	}

	if err := customValidator.ValidateName(request.FirstName); err != nil {
		return []string{err.Error()}, err
	}

	if err := customValidator.ValidateName(request.LastName); err != nil {
		return []string{err.Error()}, err
	}

	return []string{}, nil
}
