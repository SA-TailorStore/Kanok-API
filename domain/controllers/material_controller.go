package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/gofiber/fiber/v2"
)

type meterialController struct {
	service services.MaterialUseCase
}

func NewMeterialController(service services.MaterialUseCase) rest.MaterialHandler {
	return &meterialController{
		service: service,
	}
}

// CreateMaterial implements rest.MaterialHandler.
func (m *meterialController) CreateMaterial(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetMaterialByID implements rest.MaterialHandler.
func (m *meterialController) GetMaterialByID(c *fiber.Ctx) error {
	panic("unimplemented")
}
