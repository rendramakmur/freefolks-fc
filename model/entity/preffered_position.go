package entity

import "time"

type PrefferedPosition struct {
	Id                int             `gorm:"column:pp_id;primaryKey"`
	UserInformationId int             `gorm:"column:pp_pi_id"`
	UserInformation   UserInformation `gorm:"foreignKey:UserInformationId"`
	Position          int             `gorm:"column:pp_position"`
	CreatedAt         time.Time       `gorm:"column:pp_created_at"`
	CreatedBy         int             `gorm:"column:pp_created_by"`
	UpdatedAt         time.Time       `gorm:"column:pp_updated_at"`
	UpdatedBy         int             `gorm:"column:pp_updated_by"`
}

func (PrefferedPosition) TableName() string {
	return "preffered_position"
}
