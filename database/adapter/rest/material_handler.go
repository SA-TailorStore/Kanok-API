package rest

import (
	"github.com/gofiber/fiber/v2"
)

type MaterialHandler interface {
	AddMaterial(c *fiber.Ctx) error
	UpdateMaterial(c *fiber.Ctx) error
	DeleteMaterial(c *fiber.Ctx) error
	GetAllMaterials(c *fiber.Ctx) error
	GetMaterialByID(c *fiber.Ctx) error
}
