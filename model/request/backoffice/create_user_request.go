package backoffice

import (
	"github.com/rendramakmur/freefolks-fc/model/support"
)

type CreateUserRequest struct {
	Email        *string             `json:"email" validate:"required,email"`
	Password     *string             `json:"password" validate:"required"`
	UserType     *int                `json:"userType" validate:"required"`
	FirstName    *string             `json:"firstName" validate:"required"`
	LastName     *string             `json:"lastName" validate:"required"`
	MobileNumber *string             `json:"mobileNumber" validate:"required"`
	Occupation   support.DefaultData `json:"occupation" validate:"required"`
	DateOfBirth  support.OnlyDate    `json:"dob" validate:"required"`
	Gender       support.DefaultData `json:"gender" validate:"required"`
	PhotoProfile *string             `json:"photoProfile"`
	Address      *string             `json:"address"`
	City         *string             `json:"city"`
	PostalCode   *string             `json:"postalCode"`
	BodySize     support.DefaultData `json:"bodySize" validate:"required"`
}
