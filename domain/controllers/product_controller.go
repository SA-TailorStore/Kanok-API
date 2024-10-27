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

	res, err := p.service.CreateProduct(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrDesignNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrFabricNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrFailedProduct:
			return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
				"status":  "501",
				"message": "Product Create Failed",
				"data":    res,
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
		"data":    res,
	})
}

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

	products, err := p.service.GetProductByOrderID(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrProductNotFound:
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
		"status":  "200",
		"message": "Produst List By " + req.Order_id,
		"data":    products,
	})
}

func (p *productController) GetProductByID(c *fiber.Ctx) error {
	var req *requests.ProductID

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := p.service.GetProductByID(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
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
		"status":  "200",
		"message": "Get Produst",
		"data":    res,
	})
}

func (p *productController) GetAllProducts(c *fiber.Ctx) error {

	res, err := p.service.GetAllProducts(c.Context())
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
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
		"status":  "200",
		"message": "Get Produsts",
		"data":    res,
	})
}

func (p *productController) UpdateProcessQuantity(c *fiber.Ctx) error {
	var req *requests.UpdateProcessQuantity

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := p.service.UpdateProcessQuantity(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrSomethingWrong:
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

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status":  "204",
		"message": "Produst Update",
	})
}
