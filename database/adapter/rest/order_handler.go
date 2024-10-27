package rest

import (
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	CreateOrder(c *fiber.Ctx) error
	GetOrderByID(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	UpdatePayment(c *fiber.Ctx) error
	UpdateTailor(c *fiber.Ctx) error
	UpdatePrice(c *fiber.Ctx) error
	UpdateTracking(c *fiber.Ctx) error
	GetOrderByJWT(c *fiber.Ctx) error
	CheckProcess(c *fiber.Ctx) error
	GetAllOrders(c *fiber.Ctx) error
}
