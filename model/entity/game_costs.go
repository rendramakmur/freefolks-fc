package entity

import "time"

type GameCosts struct {
	Id          *int       `gorm:"column:gcs_id"`
	GameDataId  *int       `gorm:"column:gcs_gd_id"`
	GameData    *GameData  `gorm:"foreignKey:GameDataId"`
	Description *string    `gorm:"column:gcs_description"`
	Cost        *float64   `gorm:"column:gcs_cost"`
	CreatedAt   *time.Time `gorm:"column:gcs_created_at"`
	CreatedBy   *int       `gorm:"column:gcs_created_by"`
	UpdatedAt   *time.Time `gorm:"column:gcs_updated_at"`
	UpdatedBy   *int       `gorm:"column:gcs_updated_by"`
}

func (GameCosts) TableName() string {
	return "game_costs"
}
