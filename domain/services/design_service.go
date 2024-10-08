package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type DesignUseCase interface {
	CreateDesign(ctx context.Context, req *requests.CreateDesign) error
	GetDesignByID(ctx context.Context, req *requests.DesignID) error
}

type designService struct {
	reposititory reposititories.DesignRepository
	config       *configs.Config
}

func NewDesignService(reposititory reposititories.DesignRepository, config *configs.Config) DesignUseCase {
	return &designService{
		reposititory: reposititory,
		config:       config,
	}
}

// CreateDesign implements DesignUseCase.
func (d *designService) CreateDesign(ctx context.Context, req *requests.CreateDesign) error {
	panic("unimplemented")
}

// GetDesignByID implements DesignUseCase.
func (d *designService) GetDesignByID(ctx context.Context, req *requests.DesignID) error {
	panic("unimplemented")
}
