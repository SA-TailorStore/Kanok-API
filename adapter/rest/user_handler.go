package rest

import (
	"github.com/SA-TailorStore/Kanok-API/exceptions"
	"github.com/SA-TailorStore/Kanok-API/requests"
	"github.com/SA-TailorStore/Kanok-API/usercases"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FindAllUser(c *fiber.Ctx) error
}

type userHandler struct {
	userService usercases.UserUseCase
}

// FindAllUser implements UserHandler.
func (u *userHandler) FindAllUser(c *fiber.Ctx) error {
	users, err := u.userService.FindAllUser(c.Context())

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
func (u *userHandler) Login(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Register implements UserHandler.
func (u *userHandler) Register(c *fiber.Ctx) error {
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
	if err := u.userService.Register(c.Context(), req); err != nil {
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

func NewUserHandler(userService usercases.UserUseCase) UserHandler {
	return &userHandler{
		userService: userService,
	}
}
