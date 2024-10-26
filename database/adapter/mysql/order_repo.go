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

	order_id := "O" + time.Now().Format("20060102") + time.Now().Format("150405")

	query := `
	SELECT
		display_name,
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
		display_name,
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
	(order_id, status,store_phone, store_address, user_phone, user_address, created_by, tailor_phone, tailor_address, tailor_id) 
	VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = o.db.QueryContext(ctx, query,
		order_id,
		"pending",
		store.Phone_number,
		store.Display_name+"|"+store.Address,
		user.Phone_number,
		user.Display_name+"|"+user.Address,
		req.Token,
		user.Phone_number,
		user.Display_name+"|"+user.Address,
		req.Token,
	)
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
		price,
		tracking_number,
		due_date,
		created_by,
		user_phone,
		user_address,
		tailor_id,
		tailor_phone, 
		tailor_address,
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
		is_payment = ?,
		status = waiting_assign
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

func (o *OrderMySQL) UpdateTailor(ctx context.Context, req *requests.UpdateTailor) error {

	if err := utils.CheckUserByID(o.db, ctx, req.Tailor_id); err != nil {
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

	var tailor requests.UserUpdate

	if err := o.db.GetContext(ctx,
		&tailor,
		"SELECT phone_number,display_name,address FROM USERS WHERE user_id = ?",
		req.Tailor_id,
	); err != nil {
		return exceptions.ErrUserNotFound
	}

	query := `
	UPDATE ORDERS 
	SET 
		tailor_id = ?,
		tailor_phone = ?,
		tailor_address = ?,
		due_date = ?
	WHERE order_id = ?
	`

	if _, err := o.db.ExecContext(ctx, query,
		req.Tailor_id,
		tailor.Phone_number,
		tailor.Display_name+"|"+tailor.Address,
		parsedDate,
		req.Order_id,
	); err != nil {
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
		status,
		timestamp
	FROM ORDERS WHERE
	`

	var role requests.UserRole
	if err := o.db.GetContext(ctx, &role,
		"SELECT role FROM USERS WHERE user_id = ?",
		req.User_id,
	); err != nil {
		return nil, err
	}

	switch role.Role {
	case "user":
		role = requests.UserRole{Role: " created_by = ?"}
	case "tailor":
		role = requests.UserRole{Role: " tailor_id = ?"}
	}

	rows, err := o.db.QueryContext(ctx, query+role.Role, req.User_id)
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

func (o *OrderMySQL) GetAllOrders(ctx context.Context) ([]*responses.ShowOrder, error) {
	query := `
	SELECT 
		order_id,
		status,
		timestamp
	FROM ORDERS
	`

	rows, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*responses.ShowOrder, 0)
	for rows.Next() {
		var order responses.ShowOrder
		if err := rows.Scan(
			&order.Order_id,
			&order.Status,
			&order.Timestamp,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (o *OrderMySQL) CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.CheckProcess, error) {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return nil, err
	}

	var res *responses.CheckProcess
	query := `
	SELECT 
		process_quantity,
		total_quantity
	FROM PRODUCTS WHERE created_by = ?
	`

	rows, err := o.db.QueryContext(ctx, query, req.Order_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var process int
	var total int
	for rows.Next() {
		var product responses.ProductProcess
		if err := rows.Scan(
			&product.Process_quantity,
			&product.Total_quantity,
		); err != nil {
			return nil, err
		}
		process += product.Process_quantity
		total += product.Total_quantity
	}

	res = &responses.CheckProcess{Is_ready: process == total}
	return res, nil
}
