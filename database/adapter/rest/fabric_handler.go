package rest

import (
	"github.com/gofiber/fiber/v2"
)

type FabricHandler interface {
	AddFabric(c *fiber.Ctx) error
	UpdateFabric(c *fiber.Ctx) error
	DeleteFabric(c *fiber.Ctx) error
	GetFabricByID(c *fiber.Ctx) error
	GetAllFabrics(c *fiber.Ctx) error
}
