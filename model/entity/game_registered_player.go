package entity

import "time"

type GameRegisteredPlayer struct {
	Id                int             `gorm:"column:grp_id;primaryKey"`
	UserInformationId int             `gorm:"column:grp_pi_id"`
	UserInformation   UserInformation `gorm:"foreginKey:UserInformationId"`
	GameDataId        int             `gorm:"column:grp_gd_id"`
	GameData          GameData        `gorm:"foreignKey:GameDataId"`
	IsOutfield        string          `gorm:"column:grp_is_outfield"`
	AmountPaid        float64         `gorm:"column:grp_amount_paid"`
	PaidAt            time.Time       `gorm:"column:grp_paid_at"`
	TransactionNumber string          `gorm:"column:grp_transaction_number"`
	CreatedAt         time.Time       `gorm:"column:grp_created_at"`
	CreatedBy         int             `gorm:"column:grp_created_by"`
	UpdatedAt         time.Time       `gorm:"column:grp_updated_at"`
	UpdatedBy         int             `gorm:"column:grp_updated_by"`
}

func (GameRegisteredPlayer) TableName() string {
	return "game_registered_player"
}
