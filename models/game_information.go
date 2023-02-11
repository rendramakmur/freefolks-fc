package models

import "time"

type GameInformation struct {
	Id          int       `gorm:"column:gi_id;primaryKey"`
	GameDataId  int       `gorm:"column:gi_gd_id"`
	GameData    GameData  `gorm:"foreignKey:GameDataId"`
	Type        int       `gorm:"column:gi_type"`
	Description string    `gorm:"column:gi_description"`
	CreatedAt   time.Time `gorm:"column:gi_created_at"`
	CreatedBy   int       `gorm:"column:gi_created_by"`
	UpdatedAt   time.Time `gorm:"column:gi_updated_at"`
	UpdatedBy   int       `gorm:"column:gi_updated_by"`
}
