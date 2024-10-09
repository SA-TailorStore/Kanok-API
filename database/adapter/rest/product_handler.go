package rest

import (
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	CreateProduct(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	GetProductByOrderID(c *fiber.Ctx) error
}
