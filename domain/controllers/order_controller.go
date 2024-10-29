package controllers

import (
	_ "image/jpeg"
	_ "image/png"

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
		case exceptions.ErrFabricNotEnough:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
				"data":   res.Products,
			})
		case exceptions.ErrFabricNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrDesignNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrFailedProduct:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrInfomation:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
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
		"data":     res.Products,
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

	res, err := o.service.GetOrderByID(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
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

func (o *orderController) UpdateStatus(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UpdateStatus

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

	err := o.service.UpdateStatus(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
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
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Order has update",
		"status":  "204",
	})
}

func (o *orderController) UpdatePayment(c *fiber.Ctx) error {
	// Parse request
	var req requests.UpdatePayment

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

	file, err := utils.OpenFile(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := o.service.UpdatePayment(c.Context(), &req, file); err != nil {
		switch err {
		case exceptions.ErrHasPayment:
			return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "304",
			})
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrWrongSlip:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrNoImage:
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
		"message": "Order has payment",
		"status":  "204",
	})
}

func (o *orderController) UpdatePrice(c *fiber.Ctx) error {
	// Parse request
	var req requests.UpdatePrice

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

	if err := o.service.UpdatePrice(c.Context(), &req); err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrPriceIsValid:
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
		"message": "Order update price",
		"status":  "204",
	})
}

func (o *orderController) UpdateTracking(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UpdateTracking

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

	err := o.service.UpdateTracking(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
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

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Order has update",
		"status":  "204",
	})
}

func (o *orderController) UpdateTailor(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UpdateTailor

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := o.service.UpdateTailor(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrDateInvalid:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrDateToLow:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrUserNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrInvalidToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		case exceptions.ErrExpiredToken:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "401",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "500",
			})
		}
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Assign",
		"status":  "204",
	})
}

func (o *orderController) GetOrderByJWT(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserJWT

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

	res, err := o.service.GetOrderByJWT(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrExpiredToken:
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

func (o *orderController) GetAllOrders(c *fiber.Ctx) error {
	res, err := o.service.GetAllOrders(c.Context())
	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
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
		"message": "Get Orders",
		"status":  "200",
		"data":    res,
	})
}

func (o *orderController) CheckProcess(c *fiber.Ctx) error {
	// Parse request
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

	res, err := o.service.CheckProcess(c.Context(), req)
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
		"message": res,
	})
}
