package rest

import (
	"github.com/gofiber/fiber/v2"
)

type FabricHandler interface {
	CreateFabric(c *fiber.Ctx) error
	GetFabricByID(c *fiber.Ctx) error
}
