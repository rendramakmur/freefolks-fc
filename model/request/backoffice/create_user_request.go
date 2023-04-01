package backoffice

import (
	"github.com/rendramakmur/freefolks-fc/model/support"
)

type CreateUserRequest struct {
	Email        *string          `json:"email" validate:"required,email"`
	Password     *string          `json:"password" validate:"required"`
	UserType     *int             `json:"userType" validate:"required"`
	FirstName    *string          `json:"firstName" validate:"required"`
	LastName     *string          `json:"lastName" validate:"required"`
	MobileNumber *string          `json:"mobileNumber" validate:"required"`
	Occupation   *int             `json:"occupation"`
	DateOfBirth  support.OnlyDate `json:"dob" validate:"required"`
	Gender       *string          `json:"gender"`
	PhotoProfile *string          `json:"photoProfile"`
	Address      *string          `json:"address"`
	City         *string          `json:"city"`
	PostalCode   *string          `json:"postalCode"`
	BodySize     *string          `json:"bodySize"`
}
