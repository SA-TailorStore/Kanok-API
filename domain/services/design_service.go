package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type DesignUseCase interface {
	AddDesign(ctx context.Context, req *requests.AddDesign) error
	UpdateDesign(ctx context.Context, req *requests.UpdateDesign) error
	DeleteDesign(ctx context.Context, req *requests.DesignID) error
	GetAllDesigns(ctx context.Context) ([]*responses.Design, error)
	GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error)
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
func (d *designService) AddDesign(ctx context.Context, req *requests.AddDesign) error {
	err := d.reposititory.AddDesign(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDesign implements DesignUseCase.
func (d *designService) UpdateDesign(ctx context.Context, req *requests.UpdateDesign) error {
	err := d.reposititory.UpdateDesign(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDesign implements DesignUseCase.
func (d *designService) DeleteDesign(ctx context.Context, req *requests.DesignID) error {
	err := d.reposititory.DeleteDesign(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// GetAllDesigns implements DesignUseCase.
func (d *designService) GetAllDesigns(ctx context.Context) ([]*responses.Design, error) {
	var designs []*responses.Design
	return designs, nil
}

// GetDesignByID implements DesignUseCase.
func (d *designService) GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error) {
	var design *responses.Design

	return design, nil
}
