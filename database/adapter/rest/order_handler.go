package rest

import (
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	CreateOrder(c *fiber.Ctx) error
	GetOrderByID(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	GetOrderByJWT(c *fiber.Ctx) error
	UpdatePayment(c *fiber.Ctx) error
}
