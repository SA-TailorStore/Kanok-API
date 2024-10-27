package utils

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/jmoiron/sqlx"
)

func CheckUserByID(db *sqlx.DB, ctx context.Context, id string) error {

	var user responses.User

	query := `
	SELECT 
		*
	FROM USERS WHERE user_id = ?`

	err := db.GetContext(ctx, &user, query, id)
	if err != nil {
		return exceptions.ErrUserNotFound
	}

	return nil
}

func CheckUsernameDup(db *sqlx.DB, ctx context.Context, username string) error {

	query := `
	SELECT 
		username
	FROM USERS WHERE username = ?`
	err := db.GetContext(ctx, &responses.Username{}, query, username)
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		return exceptions.ErrUsernameDuplicated
	}

	return nil
}

func CheckOrderPayment(db *sqlx.DB, ctx context.Context, id string) error {
	var order responses.Order
	query := `
	SELECT 
		*
	FROM ORDERS WHERE order_id = ?`
	err := db.GetContext(ctx, &order, query, id)
	if err != nil {
		return err
	}

	if order.Is_payment == 1 {
		return exceptions.ErrHasPayment
	}

	return nil
}
func CheckOrderByID(db *sqlx.DB, ctx context.Context, id string) error {
	var order responses.Order
	query := `
	SELECT 
		*
	FROM ORDERS WHERE order_id = ?`
	err := db.GetContext(ctx, &order, query, id)
	if err != nil {
		return exceptions.ErrOrderNotFound
	}

	return nil
}

func CheckProductByID(db *sqlx.DB, ctx context.Context, id string) error {

	query := `
	SELECT 
		*
	FROM PRODUCTS WHERE product_id = ?`

	if err := db.GetContext(ctx, &responses.Product{}, query, id); err != nil {
		return exceptions.ErrProductNotFound
	}

	return nil
}

func CheckDesignByID(db *sqlx.DB, ctx context.Context, id int) error {

	query := `
	SELECT 
		design_id
	FROM DESIGNS WHERE design_id = ?`

	if err := db.GetContext(ctx, &responses.DesignID{}, query, id); err != nil {
		return exceptions.ErrDesignNotFound
	}

	return nil
}

func CheckFabricByID(db *sqlx.DB, ctx context.Context, id int) error {

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

func CheckMaterialByID(db *sqlx.DB, ctx context.Context, id int) error {

	query := `
	SELECT 
		material_id
	FROM MATERIALS WHERE material_id = ?`

	if err := db.GetContext(ctx, &responses.MaterialID{}, query, id); err != nil {
		return exceptions.ErrMaterialNotFound
	}

	return nil
}
