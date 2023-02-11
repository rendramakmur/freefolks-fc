package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Read() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf("error config file: %s", err))
	}
}
