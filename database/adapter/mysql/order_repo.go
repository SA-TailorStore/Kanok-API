package mysql

import (
	"context"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/jmoiron/sqlx"
)

type OrderMySQL struct {
	db *sqlx.DB
}

func NewOrderMySQL(db *sqlx.DB) reposititories.OrderRepository {
	return &OrderMySQL{
		db: db,
	}
}

// CreateOrder implements reposititories.OrderRepository.
func (o *OrderMySQL) CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) error {
	order_id := "o" + time.Now().Format("20060102") + time.Now().Format("150405")
	_, err := o.db.QueryContext(ctx,
		"INSERT INTO ORDERS (order_id, status, is_payment, store_phone, store_address, user_phone, user_address, due_date, create_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		order_id, "-", 0, req.Store_phone, req.Store_address, req.User_phone, req.User_address, nil, req.Create_by)

	if err != nil {
		return err
	}
	return err
}

// GetOrderByID implements reposititories.OrderRepository.
func (o *OrderMySQL) GetOrderByID(ctx context.Context, req *requests.OrderIDRequest) error {
	panic("unimplemented")
}
