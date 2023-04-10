package helper

import (
	"time"
)

func GetAge(bd *time.Time) int {
	if bd == nil {
		return 0
	}

	today := time.Now()
	age := today.Year() - bd.Year()
	if today.YearDay() < bd.YearDay() {
		age--
	}
	return age
}
