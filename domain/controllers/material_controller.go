package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
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

// CreateMaterial implements rest.MaterialHandler.
func (m *materialController) CreateMaterial(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetMaterialByID implements rest.MaterialHandler.
func (m *materialController) GetMaterialByID(c *fiber.Ctx) error {
	panic("unimplemented")
}
