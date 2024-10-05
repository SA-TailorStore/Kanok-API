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

// CreateFabric implements reposititories.FabricRepository.
func (f *FabricMySQL) CreateFabric(ctx context.Context, req *requests.CreateFabricRequest) error {
	panic("unimplemented")
}
