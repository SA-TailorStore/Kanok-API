package rest

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	UserRegister(c *fiber.Ctx) error
	StoreRegister(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
	GetAllTailor(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	GetUserByJWT(c *fiber.Ctx) error
	LoginByToken(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	UpdateImage(c *fiber.Ctx) error
}
