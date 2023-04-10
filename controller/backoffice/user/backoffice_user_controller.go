package controller

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rendramakmur/freefolks-fc/helper"
	request "github.com/rendramakmur/freefolks-fc/model/request/backoffice"
	baseResponse "github.com/rendramakmur/freefolks-fc/model/response"
	response "github.com/rendramakmur/freefolks-fc/model/response/backoffice"
	"github.com/rendramakmur/freefolks-fc/repository"
	boUserService "github.com/rendramakmur/freefolks-fc/service/backoffice/user"
	customValidator "github.com/rendramakmur/freefolks-fc/service/support"
)

type BackOfficeUserController struct {
	userService    *boUserService.BackOfficeUserService
	userRepository *repository.UserRepository
}

func NewBackOfficeUserController(userService *boUserService.BackOfficeUserService, userRepository *repository.UserRepository) *BackOfficeUserController {
	return &BackOfficeUserController{userService, userRepository}
}

func (uc *BackOfficeUserController) Login(c *fiber.Ctx) error {
	loginRequest := new(request.BackOfficeLoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, err))
	}

	validate := validator.New()
	if errRequest := validate.Struct(loginRequest); errRequest != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, errRequest))
	}

	user, err := uc.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, errors.New("Invalid email or password")))
	}

	issuedAt := jwt.NewNumericDate(time.Now())
	expiresAt := jwt.NewNumericDate(time.Now().Add(3 * time.Hour))
	claims := helper.JwtClaims{
		UserId:   *user.Id,
		Email:    *user.Email,
		UserType: *user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	}

	token, err := helper.GenerateToken(&claims)
	if err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, err))
	}

	return c.JSON(baseResponse.CreateResponse(fiber.StatusOK, response.BackOfficeLoginResponse{JwtClaims: claims, AccessToken: token}, nil))
}

func (uc *BackOfficeUserController) CreateUser(c *fiber.Ctx) error {
	createUserRequest := new(request.CreateUserRequest)

	if err := c.BodyParser(createUserRequest); err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, err))
	}

	validate := validator.New()
	if errRequest := validate.Struct(createUserRequest); errRequest != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, errRequest))
	}
	if msg, err := customValidator.ValidatePassword(createUserRequest.Password); err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, msg, err))
	}

	user, err := uc.userService.CreateUser(createUserRequest)
	if err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, err))
	}

	return c.JSON(baseResponse.CreateResponse(fiber.StatusOK, user, nil))
}
