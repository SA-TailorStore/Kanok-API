package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/usercases"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service usercases.UserUseCase
}

func NewUserController(service usercases.UserUseCase) rest.UserHandler {
	return &userController{
		service: service,
	}
}

// FindAllUser implements UserHandler.
func (u *userController) FindAllUser(c *fiber.Ctx) error {
	users, err := u.service.FindAllUser(c.Context())

	if err != nil {
		// If there's an error, return a 500 status with the error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": "Failed to retrieve users",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "User found",
		"data":    users,
	})
}

// Login implements UserHandler.
func (u *userController) Login(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserLoginRequest

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Login user
	user, err := u.service.Login(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrLoginFailed:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Login failed",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User login successfully",
		"status":  "200",
		"data":    user,
	})
}

// Register implements UserHandler.
func (u *userController) Register(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserRegisterRequest

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Register user
	if err := u.service.Register(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedUsername:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Username already registered",
				"status": "400",
			})
		case exceptions.ErrInvalidPassword:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Password must be at least 8 characters long",
				"status": "400",
			})

		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"status":  "201",
		"data":    req.Username,
	})
}
