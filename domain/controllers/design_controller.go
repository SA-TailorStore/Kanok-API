package controllers

import (
	"github.com/SA-TailorStore/Kanok-API/database/adapter/rest"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
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

// CreateDesign implements rest.DesignHandler.
func (d *designController) CreateDesign(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetDesignByID implements rest.DesignHandler.
func (d *designController) GetDesignByID(c *fiber.Ctx) error {
	panic("unimplemented")
}
