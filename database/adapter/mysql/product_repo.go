package mysql

import (
	"context"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
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

	product_id := "P" + time.Now().Format("20060102") + time.Now().Format("150405")

	_, err := p.db.QueryContext(ctx,
		"INSERT INTO PRODUCTS (product_id, detail, size, max_quantity, create_by) VALUES ( ?, ?, ?, ?, ?)",
		product_id, req.Detail, req.Size, req.Max_quantity, req.Create_by)

	if err != nil {
		return err
	}

	return err
}
