package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	service services.ProductUsecase
}

func NewProductController(service services.ProductUsecase) rest.ProductHandler {
	return &productController{
		service: service,
	}
}

// CreateProduct implements rest.ProductHandler.
func (p *productController) CreateProduct(c *fiber.Ctx) error {
	var req *requests.CreateProduct

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := p.service.CreateProduct(c.Context(), req)
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "201",
		"message": "Product Create Success",
	})
}

// GetProductByOrderID implements rest.ProductHandler.
func (p *productController) GetProductByOrderID(c *fiber.Ctx) error {
	var req *requests.OrderID

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	produsts, err := p.service.GetProductByOrderID(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Not found product",
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
		"status":  "200",
		"message": "Produst List By " + req.Order_id,
		"data":    produsts,
	})
}
