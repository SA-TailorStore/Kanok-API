package rest

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FindAllUser(c *fiber.Ctx) error
	GetUserByJWT(c *fiber.Ctx) error
	LoginToken(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	UploadImage(c *fiber.Ctx) error
}
