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

	if _, err := p.db.QueryContext(ctx, query,
		product_id,
		req.Design_id,
		req.Fabric_id,
		req.Detail,
		req.Size,
		req.Total_quantity,
		order_id,
	); err != nil {
		return err
	}

	// Update Fabric
	if _, err := p.db.ExecContext(ctx, `
		UPDATE FABRICS 
		SET 
			quantity = quantity + ? 
		WHERE fabric_id = ?`, -req.Total_quantity, req.Fabric_id,
	); err != nil {
		return err
	}

	return nil
}

func (p *ProductMySQL) GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.Product, error) {

	if err := utils.CheckOrderByID(p.db, ctx, req.Order_id); err != nil {
		return nil, err
	}

	query := `
	SELECT
          p.product_id,
		  p.design_id,
		  p.fabric_id,
		  p.detail,
		  p.size,
		  p.process_quantity,
		  p.total_quantity,
          p.created_by,
		  p.timestamp,
		  d.design_url
     FROM PRODUCTS p
	JOIN DESIGNS d ON d.design_id = p.design_id
	WHERE p.created_by = ?;
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
	err := p.db.GetContext(ctx,
		&product,
		query,
		req.Product_id,
	)

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

func (p *ProductMySQL) GetAllProducts(ctx context.Context) ([]*responses.Product, error) {

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
	FROM PRODUCTS
	`

	rows, err := p.db.QueryContext(ctx, query)
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

func (p *ProductMySQL) UpdateProcessQuantity(ctx context.Context, req *requests.UpdateProcessQuantity) error {

	if err := utils.CheckProductByID(p.db, ctx, req.Product_id); err != nil {
		return err
	}

	var process_current responses.ProductProcess
	if err := p.db.GetContext(ctx,
		&process_current,
		`
	SELECT 
		process_quantity,
		total_quantity
	FROM PRODUCTS WHERE product_id = ?`,
		req.Product_id,
	); err != nil {
		return err
	}

	query := `
	UPDATE PRODUCTS
	SET
		process_quantity = process_quantity + ?
	WHERE product_id = ?`

	if req.Increase_quantity > 0 &&
		process_current.Process_quantity <= process_current.Total_quantity &&
		req.Increase_quantity <= (process_current.Total_quantity-process_current.Process_quantity) &&
		req.Decrease_quantity == 0 {
		_, err := p.db.ExecContext(ctx,
			query,
			req.Increase_quantity,
			req.Product_id,
		)
		return err
	} else if req.Decrease_quantity > 0 &&
		process_current.Process_quantity > 0 &&
		0 <= (process_current.Process_quantity-req.Decrease_quantity) &&
		req.Increase_quantity == 0 {
		_, err := p.db.ExecContext(ctx,
			query,
			-req.Decrease_quantity,
			req.Product_id,
		)
		return err
	} else {
		return exceptions.ErrSomethingWrong
	}
}

func (p *ProductMySQL) CheckFabric(ctx context.Context, req *requests.Product, index string) (*responses.CheckFabric, error) {
	var res *responses.CheckFabric

	if err := utils.CheckFabricByID(p.db, ctx, req.Fabric_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		quantity 
	FROM FABRICS WHERE fabric_id = ?
	`

	var fabric responses.FabricQuantity
	err := p.db.GetContext(ctx, &fabric, query, req.Fabric_id)
	if err != nil {
		return nil, err
	}
	res = &responses.CheckFabric{Product_index: index, Quantity: fabric.Quantity-req.Total_quantity >= 0}

	return res, nil
}

func (p *ProductMySQL) CheckProcess(ctx context.Context, req *requests.ProductID) (*responses.ProductProcess, error) {

	if err := utils.CheckProductByID(p.db, ctx, req.Product_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		process_quantity, 
    	total_quantity
	FROM PRODUCTS WHERE product_id = ?
	`
	var product responses.ProductProcess
	err := p.db.GetContext(ctx, &product, query, req.Product_id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
