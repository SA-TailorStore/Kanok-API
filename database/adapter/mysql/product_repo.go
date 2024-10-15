package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
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

func (p *ProductMySQL) CreateProduct(ctx context.Context, req *requests.Product, order_id string, index string) error {
	// Validate
	if err := utils.CheckDesignByID(p.db, ctx, req.Design_id); err != err {
		return err
	}

	if err := utils.CheckFabricByID(p.db, ctx, req.Fabric_id); err != err {
		return err
	}

	if err := utils.CheckOrderByID(p.db, ctx, order_id); err != nil {
		return err
	}

	// Query
	query := `INSERT INTO PRODUCTS
	(product_id, design_id, fabric_id, detail, size, total_quantity, created_by) 
	VALUES ( ?, ?, ?, ?, ?, ?, ?)`

	product_id := "P" + time.Now().Format("20060102") + time.Now().Format("150405") + index

	_, err := p.db.QueryContext(ctx, query,
		product_id,
		req.Design_id,
		req.Fabric_id,
		req.Detail,
		req.Size,
		req.Total_quantity,
		order_id)

	if err != nil {
		return err
	}

	return err
}

func (p *ProductMySQL) GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.Product, error) {

	if err := utils.CheckOrderByID(p.db, ctx, req.Order_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		product_id,
		design_id,
		fabric_id,
		detail,
		size,
		process_quantity,
		total_quantity,
		created_by,
		timestamp
	FROM PRODUCTS WHERE created_by = ?
	`

	rows, err := p.db.QueryContext(ctx, query, req.Order_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*responses.Product
	for rows.Next() {
		var product responses.Product
		if err := rows.Scan(
			&product.Product_id,
			&product.Design_id,
			&product.Fabric_id,
			&product.Detail,
			&product.Size,
			&product.Process_quantity,
			&product.Total_quantity,
			&product.Created_by,
			&product.Timestamp); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, err
}

func (p *ProductMySQL) GetProductByID(ctx context.Context, req *requests.ProductID) (*responses.Product, error) {

	if err := utils.CheckProductByID(p.db, ctx, req.Product_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		product_id,
		design_id,
		fabric_id,
		detail,
		size,
		process_quantity,
		total_quantity,
		created_by,
		timestamp
	FROM PRODUCTS WHERE product_id = ?`

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
