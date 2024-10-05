package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/jmoiron/sqlx"
)

type DesignMySQL struct {
	db *sqlx.DB
}

func NewDesignMySQL(db *sqlx.DB) reposititories.DesignRepository {
	return &DesignMySQL{
		db: db,
	}
}

// CreateDesign implements reposititories.DesignRepository.
func (d *DesignMySQL) CreateDesign(ctx context.Context, req *requests.CreateOrderRequest) error {
	panic("unimplemented")
}

// GetDesignByID implements reposititories.DesignRepository.
func (d *DesignMySQL) GetDesignByID(ctx context.Context, req *requests.OrderIDRequest) error {
	panic("unimplemented")
}
