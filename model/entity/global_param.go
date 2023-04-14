package entity

import "time"

type GlobalParam struct {
	Id          *int       `gorm:"column:gp_id;primaryKey"`
	CodeId      *int       `gorm:"column:gp_code_id"`
	Slug        *string    `gorm:"column:gp_slug"`
	Name        *string    `gorm:"column:gp_name"`
	Description *string    `gorm:"column:gp_description"`
	CreatedAt   *time.Time `gorm:"column:gp_created_at"`
	UpdatedAt   *time.Time `gorm:"column:gp_updated_at"`
}

func (GlobalParam) TableName() string {
	return "global_param"
}
