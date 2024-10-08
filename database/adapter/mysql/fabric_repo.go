package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/jmoiron/sqlx"
)

type FabricMySQL struct {
	db *sqlx.DB
}

func NewFabricMySQL(db *sqlx.DB) reposititories.FabricRepository {
	return &FabricMySQL{
		db: db,
	}
}

// AddFabric implements reposititories.FabricRepository.
func (f *FabricMySQL) AddFabric(ctx context.Context, req *requests.AddFabric) error {
	panic("unimplemented")
}

// DeleteFabric implements reposititories.FabricRepository.
func (f *FabricMySQL) DeleteFabric(ctx context.Context, req *requests.FabricID) error {
	panic("unimplemented")
}

// GetAllFabrics implements reposititories.FabricRepository.
func (f *FabricMySQL) GetAllFabrics(ctx context.Context) error {
	panic("unimplemented")
}

// GetFabricByID implements reposititories.FabricRepository.
func (f *FabricMySQL) GetFabricByID(ctx context.Context, req *requests.FabricID) error {
	panic("unimplemented")
}

// UpdateFabric implements reposititories.FabricRepository.
func (f *FabricMySQL) UpdateFabric(ctx context.Context, req *requests.UpdateFabric) error {
	panic("unimplemented")
}
