package rest

import (
	"github.com/gofiber/fiber/v2"
)

type DesignHandler interface {
	AddDesign(c *fiber.Ctx) error
	UpdateDesign(c *fiber.Ctx) error
	DeleteDesign(c *fiber.Ctx) error
	GetAllDesigns(c *fiber.Ctx) error
	GetDesignByID(c *fiber.Ctx) error
}
