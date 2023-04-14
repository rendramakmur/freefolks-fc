package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rendramakmur/freefolks-fc/config"
	boUserController "github.com/rendramakmur/freefolks-fc/controller/backoffice/user"
	"github.com/rendramakmur/freefolks-fc/repository"
	boUserRoute "github.com/rendramakmur/freefolks-fc/route/backoffice/user"
	boUserService "github.com/rendramakmur/freefolks-fc/service/backoffice/user"
)

func main() {
	config.Read()
	var db = config.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Welcome to Freefolks FC")
	})

	userRepository := repository.NewUserRepository(db)
	globalParamRepo := repository.NewGlobalParamRepository(db)
	backOfficeUserService := boUserService.NewBackOfficeUserService(userRepository, globalParamRepo)
	backOfficeUserController := boUserController.NewBackOfficeUserController(backOfficeUserService, userRepository, globalParamRepo)
	boUserRoute.NewBackOfficeUserRoutes(app, backOfficeUserController)

	if err := app.Listen(os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
