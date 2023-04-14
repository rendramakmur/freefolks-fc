package user

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
)

type BackOfficeUserController struct {
	userService           *boUserService.BackOfficeUserService
	userRepository        *repository.UserRepository
	globalParamRepository *repository.GlobalParamRepository
}

func NewBackOfficeUserController(userService *boUserService.BackOfficeUserService, userRepository *repository.UserRepository, globalParamRepository *repository.GlobalParamRepository) *BackOfficeUserController {
	return &BackOfficeUserController{userService, userRepository, globalParamRepository}
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

	if msg, err := boUserService.ValidateCreateUserRequest(createUserRequest, *uc.userRepository, *uc.globalParamRepository); err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, msg, err))
	}

	user, err := uc.userService.CreateUser(createUserRequest)
	if err != nil {
		return c.JSON(baseResponse.CreateResponse(fiber.ErrBadRequest.Code, nil, err))
	}

	return c.JSON(baseResponse.CreateResponse(fiber.StatusOK, user, nil))
}
