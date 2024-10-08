package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
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
func (o *OrderMySQL) CreateOrder(ctx context.Context, req *requests.CreateOrder) error {
	order_id := "O" + time.Now().Format("20060102") + time.Now().Format("150405")
	_, err := o.db.QueryContext(ctx,
		"INSERT INTO ORDERS (order_id, is_payment, store_phone, store_address, user_phone, user_address, due_date, create_by) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)",
		order_id, 0, req.Store_phone, req.Store_address, req.User_phone, req.User_address, nil, req.Create_by)

	fmt.Println(err)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return exceptions.ErrInfomation
		default:
			return err
		}
	}
	return err
}

// GetOrderByID implements reposititories.OrderRepository.
func (o *OrderMySQL) GetOrderByID(ctx context.Context, req *requests.OrderID) error {
	panic("unimplemented")
}
