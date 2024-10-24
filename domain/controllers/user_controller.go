package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service services.UserUseCase
}

func NewUserController(service services.UserUseCase) rest.UserHandler {
	return &userController{
		service: service,
	}
}

func (u *userController) GetAllUser(c *fiber.Ctx) error {
	// Parse request
	var req requests.UserRole

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := u.service.GetAllUser(c.Context(), &req)

	if err != nil {
		switch err {
		case exceptions.ErrRoleNotHave:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "500",
				"message": "Failed to retrieve users",
				"error":   err.Error(),
			})
		}
	}

	if res != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Found:" + req.Role,
			"data":    res,
		})

	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Not Found: " + req.Role,
		})
	}
}

func (u *userController) GetAllTailor(c *fiber.Ctx) error {

	req := requests.UserRole{Role: "tailor"}

	res, err := u.service.GetAllUser(c.Context(), &req)

	if err != nil {
		switch err {
		case exceptions.ErrRoleNotHave:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "500",
				"message": "Failed to retrieve users",
				"error":   err.Error(),
			})
		}
	}

	if res != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Found:" + req.Role,
			"data":    res,
		})

	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Not Found: " + req.Role,
		})
	}

}

func (u *userController) UserRegister(c *fiber.Ctx) error {

	// Parse request
	var req *requests.UserRegister

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
		case exceptions.ErrUsernameDuplicated:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrCharLeastPassword:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrOneSpecialPassword:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrPhoneNumber:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"status":  "201",
	})
}

func (u *userController) StoreRegister(c *fiber.Ctx) error {

	// Parse request
	var req *requests.UserRegister

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
	if err := u.service.StoreRegister(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrUsernameDuplicated:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrCharLeastPassword:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrOneSpecialPassword:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrPhoneNumber:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"status":  "201",
	})
}

func (u *userController) Login(c *fiber.Ctx) error {

	// Parse request
	var req *requests.UserLogin

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Login res
	res, err := u.service.Login(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrWrongUsername:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrWrongPassword:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrLoginFailed:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"status":  "200",
		"token":   res.Token,
	})
}

func (u *userController) GetUserByJWT(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserJWT

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := u.service.GetByJWT(c.Context(), req)

	if err != nil {
		switch err {
		case exceptions.ErrInvalidToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OK!",
		"status":  "200",
		"data":    res,
	})
}

func (u *userController) StoreAssign(c *fiber.Ctx) error {
	// Parse request
	var req *requests.StoreAssign

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := u.service.StoreAssign(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrDateInvalid:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrUserNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrInvalidToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Assign",
		"status":  "204",
	})
}

func (u *userController) LoginByToken(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserJWT

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := u.service.GenerateToken(c.Context(), req)

	if err != nil {
		switch err {
		case exceptions.ErrInvalidToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "The session is still alive.",
		"status":  "201",
		"token":   res.Token,
	})
}

func (u *userController) UpdateAddress(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserUpdate

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := u.service.UpdateAddress(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrUserNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrInvalidToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}

	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Update Address success.",
		"status":  "204",
	})
}

func (u *userController) UpdateImage(c *fiber.Ctx) error {
	// Parse request
	var req requests.UserUploadImage

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	file, err := utils.OpenFile(c)
	if err != nil {
		return err
	}

	res, err := u.service.UploadImage(c.Context(), file, &req)
	if err != nil {
		switch err {
		case exceptions.ErrUploadImage:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrInvalidToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"data":    res.User_profile_url,
		"message": "Upload Image success",
		"status":  "204",
	})
}
