package utils

import (
	"context"
	"fmt"

	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/jmoiron/sqlx"
)

func CheckUser(db *sqlx.DB, ctx context.Context, id string) error {

	query := `
	SELECT 
		user_id
	FROM USERS WHERE user_id = ?`

	if err := db.GetContext(ctx, &responses.UserID{}, query, id); err != nil {
		return exceptions.ErrUserNotFound
	}

	return nil
}

func CheckOrder(db *sqlx.DB, ctx context.Context, id string) error {

	query := `
	SELECT 
		order_id
	FROM ORDERS WHERE order_id = ?`

	if err := db.GetContext(ctx, &responses.OrderID{}, query, id); err != nil {
		return exceptions.ErrOrderNotFound
	}

	return nil
}

func CheckProduct(db *sqlx.DB, ctx context.Context, id string) error {

	query := `
	SELECT 
		product_id
	FROM PRODUCTS WHERE product_id = ?`

	if err := db.GetContext(ctx, &responses.ProductID{}, query, id); err != nil {
		return exceptions.ErrProductNotFound
	}

	return nil
}

func CheckDesign(db *sqlx.DB, ctx context.Context, id int) error {

	query := `
	SELECT 
		design_id
	FROM DESIGNS WHERE design_id = ?`

	if err := db.GetContext(ctx, &responses.DesignID{}, query, id); err != nil {
		return exceptions.ErrDesignNotFound
	}

	return nil
}

func CheckFabric(db *sqlx.DB, ctx context.Context, id int) error {

	query := `
	SELECT 
		fabric_id
	FROM FABRICS WHERE fabric_id = ?`

	if err := db.GetContext(ctx, &responses.FabricID{}, query, id); err != nil {
		fmt.Println(err)
		return exceptions.ErrFabricNotFound
	}

	return nil
}

func CheckMaterial(db *sqlx.DB, ctx context.Context, id int) error {

	query := `
	SELECT 
		material_id
	FROM MATERIALS WHERE material_id = ?`

	if err := db.GetContext(ctx, &responses.MaterialID{}, query, id); err != nil {
		return exceptions.ErrMaterialNotFound
	}

	return nil
}
