package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rendramakmur/freefolks-fc/controller"
	"github.com/rendramakmur/freefolks-fc/middleware"
)

func NewBackOfficeUserRoutes(app *fiber.App, userController *controller.BackOfficeUserController) {
	backOfficeUser := app.Group("/backoffice/")
	backOfficeUser.Post("/login", userController.Login)

	userManagement := backOfficeUser.Group("/user/", middleware.Auth)
	userManagement.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Success Access Get All")
	})
	userManagement.Get("/:customerNumber", func(c *fiber.Ctx) error {
		param := c.Params("customerNumber")
		return c.JSON(fmt.Sprintf("%v", param))
	})
	userManagement.Post("/", userController.CreateUser)
}
