package mysql

import (
	"context"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
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
func (p *ProductMySQL) CreateProduct(ctx context.Context, req *requests.CreateProduct) error {

	product_id := "P" + time.Now().Format("20060102") + time.Now().Format("150405")

	_, err := p.db.QueryContext(ctx,
		"INSERT INTO PRODUCTS (product_id, design_id, fabric_id, detail, size, max_quantity, create_by) VALUES ( ?, ?, ?, ?, ?, ?, ?)",
		product_id, req.Design_id, req.Fabric_id, req.Detail, req.Size, req.Max_quantity, req.Create_by)

	if err != nil {
		return err
	}

	return err
}

// GetProductByOrderID implements reposititories.ProductRepository.
func (p *ProductMySQL) GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.ProductID, error) {

	rows, err := p.db.QueryContext(ctx, "SELECT product_id FROM ORDERS WHERE order_id = ?", req.Order_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*responses.ProductID

	for rows.Next() {
		var product_id *responses.ProductID
		if err := rows.Scan(&product_id); err != nil {
			return nil, err
		}

		products = append(products, product_id)
	}

	return products, err
}
