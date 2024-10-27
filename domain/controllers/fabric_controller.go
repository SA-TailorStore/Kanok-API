package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type fabricController struct {
	service services.FabricUseCase
}

func NewFabricController(service services.FabricUseCase) rest.FabricHandler {
	return &fabricController{
		service: service,
	}
}

func (f *fabricController) AddFabric(c *fiber.Ctx) error {
	var req requests.AddFabric

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "400",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	file, err := utils.OpenFile(c)
	if err != nil {
		return err
	}

	err = f.service.AddFabric(c.Context(), file, &req)
	if err != nil {
		switch err {
		case exceptions.ErrUploadImage:
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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Add Fabric Success.",
		"status":  "201",
	})
}

func (f *fabricController) UpdateFabric(c *fiber.Ctx) error {
	var req requests.UpdateFabric
	// Pull form file

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "400",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	file, err := utils.OpenFile(c)
	if err != nil {
		return err
	}

	err = f.service.UpdateFabric(c.Context(), file, &req)
	if err != nil {
		switch err {
		case exceptions.ErrUploadImage:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  err.Error(),
				"status": "400",
			})
		case exceptions.ErrFabricNotFound:
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
		"message": "Update Fabric Success.",
		"status":  "204",
	})
}

func (f *fabricController) UpdateFabrics(c *fiber.Ctx) error {
	var req []*requests.UpdateFabrics

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "400",
		})
	}

	for _, fabric := range req {
		if err := utils.ValidateStruct(fabric); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	}

	if err := f.service.UpdateFabrics(c.Context(), req); err != nil {
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

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Update Fabrics Success.",
		"status":  "204",
	})
}

func (f *fabricController) DeleteFabric(c *fiber.Ctx) error {
	var req requests.FabricID

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "400",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := f.service.DeleteFabric(c.Context(), &req); err != nil {
		switch err {
		case exceptions.ErrFabricNotFound:
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
		"message": "Delete Fabric Success.",
		"status":  "204",
	})
}

func (f *fabricController) GetAllFabrics(c *fiber.Ctx) error {

	res, err := f.service.GetAllFabrics(c.Context())
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
		"message": "Get All Fabric.",
		"status":  "201",
		"data":    res,
	})
}

func (f *fabricController) GetFabricByID(c *fiber.Ctx) error {
	var req requests.FabricID

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "400",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := f.service.GetFabricByID(c.Context(), &req)
	if err != nil {
		switch err {
		case exceptions.ErrFabricNotFound:
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
		"message": "Get Fabric",
		"data":    res,
		"status":  "200",
	})
}
