package repository

import (
	"github.com/rendramakmur/freefolks-fc/model/entity"
	"github.com/rendramakmur/freefolks-fc/model/support"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) FindByEmail(email string) (*entity.UserInformation, error) {
	user := new(entity.UserInformation)
	if err := ur.db.First(&user, "ui_email = ?", email).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindByCustomerNumber(customerNumber string) (*entity.UserInformation, error) {
	user := new(entity.UserInformation)

	if err := ur.db.First(&user, "ui_customer_number = ?", customerNumber).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Save(u *entity.UserInformation) (*entity.UserInformation, error) {
	if err := ur.db.Save(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) ExistByEmail(email string) bool {
	user := new(entity.UserInformation)
	result := ur.db.First(&user, "ui_email = ?", email)

	return result.RowsAffected > 0
}

func (ur *UserRepository) FindAllUser(page int, limit int, offset int, filters []support.Filter) (result []*entity.UserInformation, totalPages *int, totalItems *int, err error) {
	users := []*entity.UserInformation{}

	db := ur.db.Offset(offset).Limit(limit)

	for _, filter := range filters {
		if filter.Key == "email" {
			db.Where("ui_email LIKE ?", "%"+filter.Value+"%")
		}
	}

	items := db.Find(&users)
	if items.Error != nil {
		return nil, nil, nil, items.Error
	}

	var itemsCount int64
	items.Count(&itemsCount)
	itemsCountInt := int(itemsCount)

	pages := itemsCountInt / limit
	if itemsCountInt%limit != 0 {
		pages++
	}

	return users, &pages, &itemsCountInt, nil
}
