package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
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

// AddFabric implements rest.FabricHandler.
func (f *fabricController) AddFabric(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateFabric implements rest.FabricHandler.
func (f *fabricController) UpdateFabric(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteFabric implements rest.FabricHandler.
func (f *fabricController) DeleteFabric(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllFabrics implements rest.FabricHandler.
func (f *fabricController) GetAllFabrics(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetFabricByID implements rest.FabricHandler.
func (f *fabricController) GetFabricByID(c *fiber.Ctx) error {
	panic("unimplemented")
}
