package user

import (
	"github.com/rendramakmur/freefolks-fc/helper"
	"github.com/rendramakmur/freefolks-fc/model/entity"
	"github.com/rendramakmur/freefolks-fc/model/request/backoffice"
	"github.com/rendramakmur/freefolks-fc/repository"
)

type BackOfficeUserService struct {
	userRepository *repository.UserRepository
}

func NewBackOfficeUserService(userRepository *repository.UserRepository) *BackOfficeUserService {
	return &BackOfficeUserService{userRepository}
}

func (bou *BackOfficeUserService) Login(email *string, password *string) (*entity.UserInformation, error) {
	user, err := bou.userRepository.FindByEmail(*email)
	if err != nil {
		return nil, err
	}

	if err := helper.CheckHashedPassword(*password, *user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (bou *BackOfficeUserService) CreateUser(cur *backoffice.CreateUserRequest) (*entity.UserInformation, error) {
	hashedPassword, err := helper.HashPassword(*cur.Password)
	if err != nil {
		return nil, err
	}

	newUser := entity.UserInformation{
		CustomerNumber: helper.GenerateCustomerNumber(),
		Email:          cur.Email,
		Password:       hashedPassword,
		UserType:       cur.UserType,
		FirstName:      cur.FirstName,
		LastName:       cur.LastName,
		MobileNumber:   cur.MobileNumber,
		Occupation:     cur.Occupation,
		DateOfBirth:    cur.DateOfBirth.Time,
		Gender:         cur.Gender,
		PhotoProfile:   cur.PhotoProfile,
		Address:        cur.Address,
		City:           cur.City,
		PostalCode:     cur.PostalCode,
		BodySize:       cur.BodySize,
	}

	savedUser, err := bou.userRepository.Save(&newUser)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}
