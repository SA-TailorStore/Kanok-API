package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/gofiber/fiber/v2"
)

type materialController struct {
	service services.MaterialUseCase
}

func NewMaterialController(service services.MaterialUseCase) rest.MaterialHandler {
	return &materialController{
		service: service,
	}
}

func (m *materialController) AddMaterial(c *fiber.Ctx) error {
	var req requests.AddMaterial
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

	err := m.service.AddMaterial(c.Context(), &req)
	if err != nil {
		switch err {
		case fiber.ErrUnauthorized:
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
		"message": "Add Material Success.",
		"status":  "201",
	})
}

func (m *materialController) UpdateMaterial(c *fiber.Ctx) error {
	var req requests.UpdateMaterial
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

	err := m.service.UpdateMaterial(c.Context(), &req)
	if err != nil {
		switch err {
		case fiber.ErrUnauthorized:
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update Material Success.",
		"status":  "204",
	})
}

func (m *materialController) DeleteMaterial(c *fiber.Ctx) error {
	var req requests.MaterialID
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

	err := m.service.DeleteMaterial(c.Context(), &req)
	if err != nil {
		switch err {
		case fiber.ErrUnauthorized:
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete Material Success.",
		"status":  "204",
	})
}

func (m *materialController) GetAllMaterials(c *fiber.Ctx) error {

	res, err := m.service.GetAllMaterials(c.Context())
	if err != nil {
		switch err {
		case fiber.ErrUnauthorized:
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get All Material.",
		"status":  "200",
		"data":    res,
	})
}

func (m *materialController) GetMaterialByID(c *fiber.Ctx) error {
	var req requests.MaterialID
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

	res, err := m.service.GetMaterialByID(c.Context(), &req)
	if err != nil {
		switch err {
		case fiber.ErrUnauthorized:
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get Material Success.",
		"status":  "200",
		"data":    res,
	})
}
