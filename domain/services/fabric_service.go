package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type FabricUseCase interface {
	AddFabric(ctx context.Context, req *requests.AddFabric) error
	UpdateFabric(ctx context.Context, req *requests.UpdateFabric) error
	DeleteFabric(ctx context.Context, req *requests.FabricID) error
	GetFabricByID(ctx context.Context, req *requests.FabricID) (*responses.Fabric, error)
	GetAllFabrics(ctx context.Context) ([]*responses.Fabric, error)
}

type fabricService struct {
	reposititory reposititories.FabricRepository
	config       *configs.Config
}

func NewFabricService(reposititory reposititories.FabricRepository, config *configs.Config) FabricUseCase {
	return &fabricService{
		reposititory: reposititory,
		config:       config,
	}
}

// AddFabric implements FabricUseCase.
func (f *fabricService) AddFabric(ctx context.Context, req *requests.AddFabric) error {
	panic("unimplemented")
}

// UpdateFabric implements FabricUseCase.
func (f *fabricService) UpdateFabric(ctx context.Context, req *requests.UpdateFabric) error {
	panic("unimplemented")
}

// DeleteFabric implements FabricUseCase.
func (f *fabricService) DeleteFabric(ctx context.Context, req *requests.FabricID) error {
	panic("unimplemented")
}

// GetAllFabrics implements FabricUseCase.
func (f *fabricService) GetAllFabrics(ctx context.Context) ([]*responses.Fabric, error) {
	panic("unimplemented")
}

// GetFabricByID implements FabricUseCase.
func (f *fabricService) GetFabricByID(ctx context.Context, req *requests.FabricID) (*responses.Fabric, error) {
	panic("unimplemented")
}
