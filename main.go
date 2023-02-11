package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendramakmur/freefolks-fc/config"
)

func main() {
	config.Read()
	var _ = config.ConnectDB()
	f := fiber.New()

	f.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Freefolks FC service API")
	})

	if err := f.Listen(":8080"); err != nil {
		panic(err)
	}
}
