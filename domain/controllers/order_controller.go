package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	service services.OrderUseCase
}

func NewOrderController(service services.OrderUseCase) rest.OrderHandler {
	return &orderController{
		service: service,
	}
}

func (o *orderController) CreateOrder(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CreateOrder

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "500",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Create Order
	res, err := o.service.CreateOrder(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrInfomation:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Order success",
		"status":   "201",
		"order_id": res.Order_id,
	})
}

func (o *orderController) GetOrderByID(c *fiber.Ctx) error {
	// Parse request
	var req *requests.OrderID

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "500",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Create Order
	res, err := o.service.GetOrderByID(c.Context(), req)
	if err != nil {
		switch err {
		case err:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get Order",
		"status":  "200",
		"data":    res,
	})
}
