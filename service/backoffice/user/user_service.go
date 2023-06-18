package user

import (
	"math/rand"
	"time"

	"github.com/rendramakmur/freefolks-fc/helper"
	"github.com/rendramakmur/freefolks-fc/model/entity"
	"github.com/rendramakmur/freefolks-fc/model/request/backoffice"
	baseResposne "github.com/rendramakmur/freefolks-fc/model/response"
	response "github.com/rendramakmur/freefolks-fc/model/response/backoffice"
	"github.com/rendramakmur/freefolks-fc/model/support"
	"github.com/rendramakmur/freefolks-fc/repository"
)

type BackOfficeUserService struct {
	userRepository        *repository.UserRepository
	globalParamRepository *repository.GlobalParamRepository
	userDetailBuilder     *BackOfficeUserDetailBuilder
}

func NewBackOfficeUserService(userRepository *repository.UserRepository, globalParamRepository *repository.GlobalParamRepository, userDetailBuilder *BackOfficeUserDetailBuilder) *BackOfficeUserService {
	return &BackOfficeUserService{userRepository, globalParamRepository, userDetailBuilder}
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

func (bou *BackOfficeUserService) CreateUser(cur *backoffice.CreateUserRequest) (*response.GeneralUserResponse, error) {
	hashedPassword, err := helper.HashPassword(*cur.Password)
	if err != nil {
		return nil, err
	}

	defaultEmailStatus := new(bool)
	*defaultEmailStatus = false

	activationCode := new(string)
	*activationCode = generateActivationCode()

	newUser := entity.UserInformation{
		CustomerNumber: helper.GenerateCustomerNumber(),
		Email:          cur.Email,
		Password:       hashedPassword,
		UserType:       cur.UserType,
		FirstName:      cur.FirstName,
		LastName:       cur.LastName,
		MobileNumber:   cur.MobileNumber,
		Occupation:     cur.Occupation.Id,
		DateOfBirth:    cur.DateOfBirth.Time,
		Gender:         cur.Gender.Id,
		PhotoProfile:   cur.PhotoProfile,
		Address:        cur.Address,
		City:           cur.City,
		PostalCode:     cur.PostalCode,
		BodySize:       cur.BodySize.Id,
		EmailStatus:    defaultEmailStatus,
		ActivationCode: activationCode,
	}

	savedUser, err := bou.userRepository.Save(&newUser)
	if err != nil {
		return nil, err
	}

	occupationData, err := bou.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.OccupationSlug, savedUser.Occupation)
	if err != nil {
		return nil, err
	}

	genderData, err := bou.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.GenderSlug, savedUser.Gender)
	if err != nil {
		return nil, err
	}

	BodySizeData, err := bou.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.BodySizeSlug, savedUser.BodySize)
	if err != nil {
		return nil, err
	}

	return &response.GeneralUserResponse{
		Id:             savedUser.Id,
		Email:          savedUser.Email,
		UserType:       savedUser.UserType,
		FirstName:      savedUser.FirstName,
		LastName:       savedUser.LastName,
		MobileNumber:   savedUser.MobileNumber,
		Occupation:     occupationData,
		DateOfBirth:    savedUser.DateOfBirth,
		Gender:         genderData,
		PhotoProfile:   savedUser.PhotoProfile,
		Address:        savedUser.Address,
		City:           savedUser.City,
		PostalCode:     savedUser.PostalCode,
		BodySize:       BodySizeData,
		ActivationCode: savedUser.ActivationCode,
		EmailStatus:    savedUser.EmailStatus,
		VerifiedAt:     savedUser.VerifiedAt,
	}, nil
}

func (bou *BackOfficeUserService) GetUserDetail(customerNumber *string) (*response.BackOfficeUserDetailResponse, error) {
	userDetail, err := bou.userDetailBuilder.Build(customerNumber)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func (bou *BackOfficeUserService) GetAllUser(gul *backoffice.GetUserListRequest) (*baseResposne.Pagination, error) {
	users := []response.UserListResponse{}
	filters := []support.Filter{}

	offset := (gul.Page - 1) * gul.Limit
	filters = append(filters, support.Filter{Key: "email", Value: gul.Email})

	result, totalPages, itemsCount, err := bou.userRepository.FindAllUser(gul.Page, gul.Limit, offset, filters)
	if err != nil {
		return nil, err
	}

	for _, user := range result {
		users = append(users, response.UserListResponse{
			Id:          user.Id,
			Email:       user.Email,
			UserType:    user.UserType,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			EmailStatus: user.EmailStatus,
		})
	}

	converted := make([]interface{}, len(users))
	for i, u := range users {
		converted[i] = u
	}

	pagination := baseResposne.Pagination{
		Page:       gul.Page,
		Limit:      gul.Limit,
		TotalPages: *totalPages,
		TotalItems: *itemsCount,
		Items:      converted,
	}

	return &pagination, nil
}

func generateActivationCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
