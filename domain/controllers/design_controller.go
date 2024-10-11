package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type designController struct {
	service services.DesignUseCase
}

func NewDesignController(service services.DesignUseCase) rest.DesignHandler {
	return &designController{
		service: service,
	}
}

func (d *designController) AddDesign(c *fiber.Ctx) error {
	// Parse Request
	var req requests.AddDesign
	// Pull form file
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to get file",
			"message": err.Error(),
		})
	}

	// Open File
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to open file",
			"message": err.Error(),
		})
	}
	defer file.Close()

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = d.service.AddDesign(c.Context(), file, &req)
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
		"message": "Add Design Success.",
		"status":  "201",
	})
}

func (d *designController) UpdateDesign(c *fiber.Ctx) error {
	// Parse Request
	var req requests.UpdateDesign

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
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

	if err := d.service.UpdateDesign(c.Context(), file, &req); err != nil {
		switch err {
		case exceptions.ErrDesignNotFound:
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
		"message": "Update Design success.",
		"status":  "204",
	})
}

func (d *designController) DeleteDesign(c *fiber.Ctx) error {
	var req requests.DesignID
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := d.service.DeleteDesign(c.Context(), &req)
	if err != nil {
		switch err {
		case exceptions.ErrDesignNotFound:
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
		"message": "Delete Design Success.",
		"status":  "204",
	})
}

func (d *designController) GetAllDesigns(c *fiber.Ctx) error {

	designs, err := d.service.GetAllDesigns(c.Context())
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
		"message": "All Designs",
		"status":  "200",
		"data":    designs,
	})
}

func (d *designController) GetDesignByID(c *fiber.Ctx) error {
	// Parse Requset
	var req requests.DesignID

	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := d.service.GetDesignByID(c.Context(), &req)
	if err != nil {
		switch err {
		case exceptions.ErrDesignNotFound:
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
		"message": "Get Design",
		"data":    res,
		"status":  "200",
	})
}
