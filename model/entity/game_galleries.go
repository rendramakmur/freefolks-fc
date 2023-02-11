package entity

import "time"

type GameGalleries struct {
	Id         int       `gorm:"column:gg_id;primaryKey"`
	GameDataId int       `gorm:"column:gg_gd_id"`
	GameData   GameData  `gorm:"foreignKey:GameDataId"`
	ImageUrl   string    `gorm:"column:gg_image_url"`
	AltImage   string    `gorm:"column:gg_alt_image"`
	CreatedAt  time.Time `gorm:"column:gg_created_at"`
	CreatedBy  int       `gorm:"column:gg_created_by"`
	UpdatedAt  time.Time `gorm:"column:gg_updated_at"`
	UpdatedBy  int       `gorm:"column:gg_updated_by"`
}
