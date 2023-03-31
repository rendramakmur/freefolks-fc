package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendramakmur/freefolks-fc/config"
	"github.com/rendramakmur/freefolks-fc/controller"
	"github.com/rendramakmur/freefolks-fc/repository"
	"github.com/rendramakmur/freefolks-fc/route"
	"github.com/rendramakmur/freefolks-fc/service"
)

func main() {
	config.Read()
	var db = config.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Welcome to Freefolks FC")
	})

	userRepository := repository.NewUserRepository(db)
	backOfficeUserService := service.NewBackOfficeUserService(userRepository)
	backOfficeUserController := controller.NewBackOfficeUserController(backOfficeUserService)
	route.NewBackOfficeUserRoutes(app, backOfficeUserController)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
