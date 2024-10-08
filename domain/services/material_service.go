package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type MaterialUseCase interface {
	CreateMaterial(ctx context.Context, req *requests.CreateMaterial) error
	GetMaterialByID(ctx context.Context, req *requests.MaterialID) error
}

type materialService struct {
	reposititory reposititories.MaterialRepository
	config       *configs.Config
}

func NewMaterialService(reposititory reposititories.MaterialRepository, config *configs.Config) MaterialUseCase {
	return &materialService{
		reposititory: reposititory,
		config:       config,
	}
}

// CreateMaterial implements MaterialUseCase.
func (m *materialService) CreateMaterial(ctx context.Context, req *requests.CreateMaterial) error {
	panic("unimplemented")
}

// GetMaterialByID implements MaterialUseCase.
func (m *materialService) GetMaterialByID(ctx context.Context, req *requests.MaterialID) error {
	panic("unimplemented")
}
