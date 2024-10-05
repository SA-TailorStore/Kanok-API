package rest

import (
	"github.com/gofiber/fiber/v2"
)

type MaterialHandler interface {
	CreateMaterial(c *fiber.Ctx) error
	GetMaterialByID(c *fiber.Ctx) error
}
