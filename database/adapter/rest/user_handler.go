package rest

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FindAllUser(c *fiber.Ctx) error
}
