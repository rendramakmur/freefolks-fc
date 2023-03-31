package repository

import (
	"github.com/rendramakmur/freefolks-fc/model/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) FindByEmail(email string) (*entity.UserInformation, error) {
	var user entity.UserInformation
	if err := ur.db.First(&user, "ui_email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Save(u *entity.UserInformation) (*entity.UserInformation, error) {
	if err := ur.db.Save(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) ExistByEmail(email string) bool {
	var user entity.UserInformation
	result := ur.db.First(&user, "ui_email = ?", email)

	return result.RowsAffected > 0
}
