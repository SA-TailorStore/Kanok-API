package mysql

import (
	"context"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
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

	query := `
	INSERT INTO ORDERS
	(order_id, store_phone, store_address, user_phone, user_address, created_by) 
	VALUES ( ?, ?, ?, ?, ?, ?)
	`

	order_id := "O" + time.Now().Format("20060102") + time.Now().Format("150405")
	_, err := o.db.QueryContext(ctx, query,
		order_id,
		req.Store_phone,
		req.Store_address,
		req.User_phone,
		req.User_address,
		req.Created_by)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderByID implements reposititories.OrderRepository.
func (o *OrderMySQL) GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error) {
	query := `
	SELECT 
		order_id
		is_payment,
		status,
		store_phone,
		store_address,
		user_phone,
		user_address,
		price,
		due_date,
		tracking_number,
		tailor_id,
		created_by,
		timestamp
	FROM ORDERS WHERE order_id = ?`
	var order responses.Order

	err := o.db.GetContext(ctx, &order, query, req.Order_id)
	if err != nil {
		return &order, err
	}
	return &order, nil
}
