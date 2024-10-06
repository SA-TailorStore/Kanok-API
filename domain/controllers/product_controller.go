package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
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
	var req *requests.CreateProductRequest

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
		return err
	}

	return err
}

// GetProductByOrderID implements rest.ProductHandler.
func (p *productController) GetProductByOrderID(c *fiber.Ctx) error {
	var req *requests.OrderIDRequest

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
		// If there's an error, return a 500 status with the error message
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": "Failed to retrieve products",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "Produst List By " + req.Order_id,
		"data":    produsts,
	})
}
