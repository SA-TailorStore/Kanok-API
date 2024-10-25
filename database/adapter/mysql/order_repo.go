package mysql

import (
	"context"
	"time"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
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

func (o *OrderMySQL) CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.OrderID, error) {
	query := `
	SELECT 
		address,
		phone_number
	FROM USERS WHERE role = "store"`

	var store responses.UserCreateOrder
	err := o.db.GetContext(ctx, &store, query)
	if err != nil {
		return nil, exceptions.ErrInfomation
	}

	query = `
	SELECT 
		address,
		phone_number 
	FROM USERS WHERE user_id = ?`

	var user responses.UserCreateOrder
	err = o.db.GetContext(ctx, &user, query, req.Token)
	if err != nil {
		return nil, exceptions.ErrInfomation
	}

	query = `
	INSERT INTO ORDERS
	(order_id, status,store_phone, store_address, user_phone, user_address, tailor_id, created_by) 
	VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)
	`

	order_id := "O" + time.Now().Format("20060102") + time.Now().Format("150405")
	_, err = o.db.QueryContext(ctx, query,
		order_id,
		"pending",
		store.Phone_number,
		store.Address,
		user.Phone_number,
		user.Display_name+"|"+user.Address,
		req.Token,
		req.Token)
	if err != nil {
		return nil, exceptions.ErrInfomation
	}
	return &responses.OrderID{Order_id: order_id}, nil
}

func (o *OrderMySQL) GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error) {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		order_id,
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

func (o *OrderMySQL) UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	query := `
	UPDATE ORDERS
	SET
		status = ?,
		price = ?
	WHERE order_id = ?`

	_, err := o.db.ExecContext(ctx, query, req.Status, req.Price, req.Order_id)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderMySQL) UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	query := `
	UPDATE ORDERS 
	SET 
		is_payment = ? 
	WHERE order_id = ?`

	_, err := o.db.ExecContext(ctx, query, req.Is_payment, req.Order_id)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderMySQL) UpdateTracking(ctx context.Context, req *requests.UpdateTracking) error {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	query := `
	UPDATE ORDERS 
	SET 
		tracking_number = ?
	WHERE order_id = ?`

	_, err := o.db.ExecContext(ctx, query, req.Tracking_number, req.Order_id)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderMySQL) GetOrderByUserId(ctx context.Context, req *requests.UserID) ([]*responses.Order, error) {

	if err := utils.CheckUserByID(o.db, ctx, req.User_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		order_id,
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
	FROM ORDERS WHERE created_by = ?
	`

	rows, err := o.db.QueryContext(ctx, query, req.User_id)
	if err != nil {
		return nil, exceptions.ErrUserNotFound
	}
	defer rows.Close()

	orders := make([]*responses.Order, 0)
	for rows.Next() {
		var order responses.Order
		if err := rows.Scan(
			&order.Order_id,
			&order.Is_payment,
			&order.Status,
			&order.Store_phone,
			&order.Store_address,
			&order.User_phone,
			&order.User_address,
			&order.Price,
			&order.Due_date,
			&order.Tracking_number,
			&order.Tailor_id,
			&order.Created_by,
			&order.Timestamp,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (o *OrderMySQL) GetAllOrders(ctx context.Context) ([]*responses.Order, error) {
	query := `
	SELECT 
		order_id,
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
	FROM ORDERS
	`

	rows, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*responses.Order, 0)
	for rows.Next() {
		var order responses.Order
		if err := rows.Scan(
			&order.Order_id,
			&order.Is_payment,
			&order.Status,
			&order.Store_phone,
			&order.Store_address,
			&order.User_phone,
			&order.User_address,
			&order.Price,
			&order.Due_date,
			&order.Tracking_number,
			&order.Tailor_id,
			&order.Created_by,
			&order.Timestamp,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (o *OrderMySQL) StoreAssign(ctx context.Context, req *requests.StoreAssign) error {

	if err := utils.CheckUserByID(o.db, ctx, req.User_id); err != nil {
		return err
	}

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	layout := time.RFC3339
	parsedDate, err := time.Parse(layout, req.Due_date)

	if err != nil {
		return exceptions.ErrDateInvalid
	}

	query := `
	UPDATE ORDERS 
	SET 
		tailor_id = ?,
		due_date = ?
	WHERE order_id = ?
	`

	if _, err := o.db.ExecContext(ctx, query, req.User_id, parsedDate, req.Order_id); err != nil {
		return err
	}

	return nil
}
