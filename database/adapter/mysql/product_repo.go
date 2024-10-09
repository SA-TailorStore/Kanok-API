package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
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
	// Validate
	query := `SELECT design_id FROM DESIGNS WHERE design_id = ?`
	err := p.db.GetContext(ctx, &responses.DesignID{}, query, req.Design_id)

	if err != nil {
		return exceptions.ErrDesignNotFound
	}

	query = `SELECT fabric_id FROM FABRICS WHERE fabric_id = ?`
	err = p.db.GetContext(ctx, &responses.FabricID{}, query, req.Fabric_id)
	if err != nil {
		return exceptions.ErrFabricNotFound
	}

	query = `SELECT order_id FROM ORDERS WHERE order_id = ?`
	err = p.db.GetContext(ctx, &responses.OrderID{}, query, req.Created_by)
	if err != nil {
		return exceptions.ErrOrderNotFound
	}

	// Query
	query = `INSERT INTO PRODUCTS
	(product_id, design_id, fabric_id, detail, size, total_quantity, created_by) 
	VALUES ( ?, ?, ?, ?, ?, ?, ?)`

	product_id := "P" + time.Now().Format("20060102") + time.Now().Format("150405")

	_, err = p.db.QueryContext(ctx, query,
		product_id,
		req.Design_id,
		req.Fabric_id,
		req.Detail,
		req.Size,
		req.Total_quantity,
		req.Created_by)

	if err != nil {
		return err
	}

	return err
}

// GetProductByOrderID implements reposititories.ProductRepository.
func (p *ProductMySQL) GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.ProductID, error) {
	query := `SELECT product_id FROM PRODUCTS WHERE created_by = ?`
	rows, err := p.db.QueryContext(ctx, query, req.Order_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*responses.ProductID

	for rows.Next() {
		var product_id responses.ProductID
		if err := rows.Scan(&product_id.Product_id); err != nil {
			return nil, err
		}

		products = append(products, &product_id)
	}

	return products, err
}

func (p *ProductMySQL) GetProductByID(ctx context.Context, req *requests.ProductID) (*responses.Product, error) {
	query :=
		`SELECT 
	product_id,
	design_id,
	fabric_id,
	detail,
	size,
	process_quantity,
	total_quantity,
	created_by,
	timestamp
	FROM PRODUCTS
	WHERE product_id = ?`

	var product responses.Product
	err := p.db.GetContext(ctx, &product,
		query, req.Product_id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return &product, exceptions.ErrProductNotFound
		default:
			return &product, err
		}
	}

	return &product, err
}
