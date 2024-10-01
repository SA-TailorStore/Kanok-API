package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductMySQL struct {
	db *sqlx.DB
}

func NewProductMySQL(db *sqlx.DB) reposititories.ProductRepository {
	return &ProductMySQL{
		db: db,
	}
}

// CreateProduct implements reposititories.ProductRepository.
func (p *ProductMySQL) CreateProduct(ctx context.Context, req *requests.CreateProductRequest) error {

	product_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	_, err = p.db.QueryContext(ctx,
		"INSERT INTO products (product_id, detail, size, quantity, create_by) VALUES ( ?, ?, ?, ?, ?)",
		product_id, req.Detail, req.Size, req.Quantity, req.Create_by)

	if err != nil {
		return err
	}

	return err
}
