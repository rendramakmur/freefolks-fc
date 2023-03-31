package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rendramakmur/freefolks-fc/helper"
	model "github.com/rendramakmur/freefolks-fc/model/response"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("X-Freefolks-Token")
	if token == "" {
		return c.JSON(model.CreateResponse(fiber.ErrForbidden.Code, nil, errors.New("Auth failed")))
	}

	claims, err := helper.ParseToken(token)
	if err != nil {
		return c.JSON(model.CreateResponse(fiber.ErrForbidden.Code, nil, errors.New("Auth failed")))
	}

	if claims.UserType != 1 {
		return c.JSON(model.CreateResponse(fiber.ErrForbidden.Code, nil, errors.New("Unauthorized user")))
	}

	return c.Next()
}
