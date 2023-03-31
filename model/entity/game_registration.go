package entity

import "time"

type GameRegistration struct {
	Id                int             `gorm:"column:gr_id;primaryKey"`
	UserInformationId int             `gorm:"column:gr_pi_id"`
	UserInformation   UserInformation `gorm:"foreginKey:UserInformationId"`
	GameDataId        int             `gorm:"column:gr_gd_id"`
	GameData          GameData        `gorm:"foreignKey:GameDataId"`
	IsOutfield        string          `gorm:"column:gr_is_outfield"`
	Amount            float64         `gorm:"column:gr_amount"`
	TransactionNumber string          `gorm:"column:gr_transaction_number"`
	CreatedAt         time.Time       `gorm:"column:gr_created_at"`
	CreatedBy         int             `gorm:"column:gr_created_by"`
	UpdatedAt         time.Time       `gorm:"column:gr_updated_at"`
	UpdatedBy         int             `gorm:"column:gr_updated_by"`
}

func (GameRegistration) TableName() string {
	return "game_registration"
}
