package entity

import "time"

type UserInformation struct {
	Id             int       `gorm:"column:ui_id;primaryKey"`
	UserType       int       `gorm:"column:ui_user_type"`
	CustomerNumber string    `gorm:"column:ui_customer_number"`
	FirstName      string    `gorm:"column:ui_first_name"`
	LastName       string    `gorm:"column:ui_last_name"`
	Email          string    `gorm:"column:ui_email"`
	Password       string    `gorm:"column:ui_password"`
	MobileNumber   string    `gorm:"column:ui_mobile_number"`
	Occupation     int       `gorm:"column:ui_occupation"`
	DateOfBirth    time.Time `gorm:"column:ui_date_of_birth"`
	Gender         string    `gorm:"column:ui_gender"`
	PhotoProfile   string    `gorm:"column:ui_photo_profile"`
	Address        string    `gorm:"column:ui_address"`
	City           string    `gorm:"column:ui_city"`
	PostalCode     string    `gorm:"column:ui_postal_code"`
	BodySize       string    `gorm:"column:ui_body_size"`
	ActivationCode string    `gorm:"column:ui_activation_code"`
	EmailStatus    string    `gorm:"column:ui_email_status"`
	VerifiedAt     time.Time `gorm:"column:ui_verified_at"`
	CreatedAt      time.Time `gorm:"column:ui_created_at"`
	CreatedBy      int       `gorm:"column:ui_created_by"`
	UpdatedAt      time.Time `gorm:"column:ui_updated_at"`
	UpdatedBy      int       `gorm:"column:ui_updated_by"`
}
