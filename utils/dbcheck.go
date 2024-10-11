package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	db *sqlx.DB
}

func NewMySQL() *sqlx.DB {
	cfg := configs.NewConfig()
	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db
}

func (db *MySQL) CheckOrder(ctx context.Context, id string) error {
	query := `
	SELECT 
		order_id
	FROM ORDERS WHERE order_id = ?`
	err := db.db.GetContext(ctx, &responses.OrderID{}, query, id)
	if err != nil {
		return exceptions.ErrOrderNotFound
	}

	return nil
}
