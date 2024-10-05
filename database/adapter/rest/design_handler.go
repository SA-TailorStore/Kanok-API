package rest

import (
	"github.com/gofiber/fiber/v2"
)

type DesignHandler interface {
	CreateDesign(c *fiber.Ctx) error
	GetDesignByID(c *fiber.Ctx) error
}
