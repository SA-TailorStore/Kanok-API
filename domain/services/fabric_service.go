package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type FabricUseCase interface {
	CreateFabric(ctx context.Context, req *requests.CreateFabricRequest) error
	GetFabricByID(ctx context.Context, req *requests.FabricIDRequest) error
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

// CreateFabric implements FabricUseCase.
func (f *fabricService) CreateFabric(ctx context.Context, req *requests.CreateFabricRequest) error {
	panic("unimplemented")
}

// GetFabricByID implements FabricUseCase.
func (f *fabricService) GetFabricByID(ctx context.Context, req *requests.FabricIDRequest) error {
	panic("unimplemented")
}
