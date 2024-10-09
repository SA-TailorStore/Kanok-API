package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type MaterialUseCase interface {
	AddMaterial(ctx context.Context, req *requests.AddMaterial) error
	UpdateMaterial(ctx context.Context, req *requests.UpdateMaterial) error
	DeleteMaterial(ctx context.Context, req *requests.MaterialID) error
	GetAllMaterials(ctx context.Context) ([]*responses.Material, error)
	GetMaterialByID(ctx context.Context, req *requests.MaterialID) (*responses.Material, error)
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

// AddMaterial implements MaterialUseCase.
func (m *materialService) AddMaterial(ctx context.Context, req *requests.AddMaterial) error {
	err := m.reposititory.AddMaterial(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMaterial implements MaterialUseCase.
func (m *materialService) UpdateMaterial(ctx context.Context, req *requests.UpdateMaterial) error {
	err := m.reposititory.UpdateMaterial(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMaterial implements MaterialUseCase.
func (m *materialService) DeleteMaterial(ctx context.Context, req *requests.MaterialID) error {
	err := m.reposititory.DeleteMaterial(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// GetAllMaterials implements MaterialUseCase.
func (m *materialService) GetAllMaterials(ctx context.Context) ([]*responses.Material, error) {
	res, err := m.reposititory.GetAllMaterials(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetMaterialByID implements MaterialUseCase.
func (m *materialService) GetMaterialByID(ctx context.Context, req *requests.MaterialID) (*responses.Material, error) {
	res, err := m.reposititory.GetMaterialByID(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}
