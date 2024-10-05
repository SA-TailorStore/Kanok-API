package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/jmoiron/sqlx"
)

type MaterialMySQL struct {
	db *sqlx.DB
}

func NewMaterialMySQL(db *sqlx.DB) reposititories.MaterialRepository {
	return &MaterialMySQL{
		db: db,
	}
}

// CreateMaterial implements reposititories.MaterialRepository.
func (m *MaterialMySQL) CreateMaterial(ctx context.Context, req *requests.CreateMaterialRequest) error {
	panic("unimplemented")
}
