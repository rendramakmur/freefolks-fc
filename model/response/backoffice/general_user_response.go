package backoffice

import (
	"time"

	"github.com/rendramakmur/freefolks-fc/model/support"
)

type GeneralUserResponse struct {
	Id             *int                `json:"id"`
	Email          *string             `json:"email"`
	UserType       *int                `json:"userType"`
	FirstName      *string             `json:"firstName"`
	LastName       *string             `json:"lastName"`
	MobileNumber   *string             `json:"mobileNumber"`
	Occupation     support.DefaultData `json:"occupation"`
	DateOfBirth    *time.Time          `json:"dateOfBirth"`
	Gender         support.DefaultData `json:"gender"`
	PhotoProfile   *string             `json:"photoProfile"`
	Address        *string             `json:"address"`
	City           *string             `json:"city"`
	PostalCode     *string             `json:"postalCode"`
	BodySize       support.DefaultData `json:"bodySize"`
	ActivationCode *string             `json:"activationCode"`
	EmailStatus    *bool               `json:"emailStatus"`
	VerifiedAt     *time.Time          `json:"verifiedAt"`
}
