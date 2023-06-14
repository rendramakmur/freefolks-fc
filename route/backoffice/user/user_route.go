package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	boUserController "github.com/rendramakmur/freefolks-fc/controller/backoffice/user"
	"github.com/rendramakmur/freefolks-fc/middleware"
)

func NewBackOfficeUserRoutes(app *fiber.App, userController *boUserController.BackOfficeUserController) {
	backOfficeUser := app.Group("/backoffice/")
	backOfficeUser.Post("/login", userController.Login)

	userManagement := backOfficeUser.Group("/user/", middleware.Auth)
	userManagement.Get("/", userController.GetCustomerList)
	userManagement.Get("/:customerNumber", func(c *fiber.Ctx) error {
		param := c.Params("customerNumber")
		return c.JSON(fmt.Sprintf("%v", param))
	})
	userManagement.Post("/", userController.CreateUser)
}
