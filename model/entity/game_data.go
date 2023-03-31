package entity

import "time"

type GameData struct {
	Id              int       `gorm:"column:gd_id;primaryKey"`
	GameNumber      string    `gorm:"column:gd_game_number"`
	VenueName       string    `gorm:"column:gd_venue_name"`
	VenueAddress    string    `gorm:"column:gd_venue_address"`
	MapUrl          string    `gorm:"column:gd_map_url"`
	Time            string    `gorm:"column:gd_time"`
	GoalkeeperQuota int       `gorm:"column:gd_goalkeeper_quota"`
	OutfieldQuota   int       `gorm:"column:gd_outfield_quota"`
	GoalkeeperPrice float64   `gorm:"column:gd_goalkeper_price"`
	OutfieldPrice   float64   `gorm:"column:gd_outfield_price"`
	TotalCost       float64   `gorm:"column:gd_total_cost"`
	Notes           string    `gorm:"column:gd_notes"`
	Status          string    `gorm:"column:gd_status"`
	CreatedAt       time.Time `gorm:"column:gd_created_at"`
	CreatedBy       int       `gorm:"column:gd_created_by"`
	UpdatedAt       time.Time `gorm:"column:gd_updated_at"`
	UpdatedBy       int       `gorm:"column:gd_updated_by"`
}

func (GameData) TableName() string {
	return "game_data"
}
