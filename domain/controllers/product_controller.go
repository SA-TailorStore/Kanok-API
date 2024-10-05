package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/gofiber/fiber/v2"
)

type productController struct {
	service services.ProductUsecase
}

// CreateProduct implements rest.ProductHandler.
func (p *productController) CreateProduct(c *fiber.Ctx) error {
	var req *requests.CreateProductRequest

	err := p.service.CreateProduct(c.Context(), req)
	if err != nil {
		return err
	}

	return err
}

// GetProductByOrderID implements rest.ProductHandler.
func (p *productController) GetProductByOrderID(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewProductController(service services.ProductUsecase) rest.ProductHandler {
	return &productController{
		service: service,
	}
}
