package user

import (
	"github.com/rendramakmur/freefolks-fc/helper"
	backofficeResponse "github.com/rendramakmur/freefolks-fc/model/response/backoffice"
	"github.com/rendramakmur/freefolks-fc/repository"
)

type BackOfficeUserDetailBuilder struct {
	userRepository        *repository.UserRepository
	globalParamRepository *repository.GlobalParamRepository
}

func NewBackOfficeUserDetailBuilder(userRepository *repository.UserRepository, globalParamRepository *repository.GlobalParamRepository) *BackOfficeUserDetailBuilder {
	return &BackOfficeUserDetailBuilder{userRepository, globalParamRepository}
}

func (udb *BackOfficeUserDetailBuilder) Build(customerNumber *string) (*backofficeResponse.BackOfficeUserDetailResponse, error) {
	userDetail, err := udb.buildUserDetail(customerNumber)
	if err != nil {
		return nil, err
	}

	return &backofficeResponse.BackOfficeUserDetailResponse{UserDetail: *userDetail}, nil
}

func (udb *BackOfficeUserDetailBuilder) buildUserDetail(customerNumber *string) (*backofficeResponse.UserDetail, error) {
	user, err := udb.userRepository.FindByCustomerNumber(*customerNumber)
	if err != nil {
		return nil, err
	}

	userTypeData, err := udb.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.UserTypeSlug, user.UserType)
	if err != nil {
		return nil, err
	}

	occupationData, err := udb.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.OccupationSlug, user.Occupation)
	if err != nil {
		return nil, err
	}

	genderData, err := udb.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.GenderSlug, user.Gender)
	if err != nil {
		return nil, err
	}

	BodySizeData, err := udb.globalParamRepository.GetDefaultDataBySlugAndCodeId(helper.BodySizeSlug, user.BodySize)
	if err != nil {
		return nil, err
	}

	userDetail := backofficeResponse.UserDetail{
		Id:             user.Id,
		CustomerNumber: user.CustomerNumber,
		Email:          user.Email,
		UserType:       userTypeData,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		MobileNumber:   user.MobileNumber,
		Occupation:     occupationData,
		DateOfBirth:    user.DateOfBirth,
		Gender:         genderData,
		PhotoProfile:   user.PhotoProfile,
		Address:        user.Address,
		City:           user.City,
		PostalCode:     user.PostalCode,
		BodySize:       BodySizeData,
		ActivationCode: user.ActivationCode,
		EmailStatus:    user.EmailStatus,
		VerifiedAt:     user.VerifiedAt,
	}

	return &userDetail, nil
}
