package mysql

import (
	"context"
	"strconv"
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

func (o *OrderMySQL) CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.CreateOrder, error) {

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

	//INSERT ORDERS
	query = `
	INSERT INTO ORDERS
	(order_id, status,store_phone, store_address, user_phone, user_address, created_by, tailor_phone, tailor_address, tailor_id) 
	VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	if _, err := o.db.QueryContext(ctx, query,
		order_id,
		"pending",
		store.Phone_number,
		store.Display_name+"|"+store.Address,
		user.Phone_number,
		user.Display_name+"|"+user.Address,
		req.Token,
		"-",
		"-",
		req.Token,
	); err != nil {
		return nil, err
	}

	// Create Product
	var fabrics = make(map[string]bool)
	var is_enough = true

	for index, value := range req.Products {
		fabric, err := o.CheckFabric(ctx, &value, strconv.Itoa(index+1))
		if err != nil {
			return nil, err
		}

		if !fabric.Quantity {
			is_enough = false
		}
		fabrics[fabric.Product_index] = fabric.Quantity
	}

	// Create Product
	if is_enough {
		for index, value := range req.Products {
			// Validate
			if err := utils.CheckDesignByID(o.db, ctx, value.Design_id); err != err {
				return nil, err
			}

			if err := utils.CheckFabricByID(o.db, ctx, value.Fabric_id); err != err {
				return nil, err
			}

			// Query
			query := `INSERT INTO PRODUCTS
			(product_id, design_id, fabric_id, detail, size, total_quantity, created_by) 
			VALUES ( ?, ?, ?, ?, ?, ?, ?)`

			product_id := "P" + time.Now().Format("20060102") + time.Now().Format("150405") + strconv.Itoa(index+1)

			if _, err := o.db.QueryContext(ctx, query,
				product_id,
				value.Design_id,
				value.Fabric_id,
				value.Detail,
				value.Size,
				value.Total_quantity,
				order_id,
			); err != nil {
				return nil, exceptions.ErrFailedProduct
			}

			// Update Fabric
			if _, err := o.db.ExecContext(ctx, `
			UPDATE FABRICS 
			SET 
				quantity = quantity + ? 
			WHERE fabric_id = ?`, -value.Total_quantity, value.Fabric_id,
			); err != nil {
				return nil, err
			}
		}
	} else {
		return &responses.CreateOrder{Order_id: order_id, Products: fabrics}, exceptions.ErrFabricNotEnough
	}

	return &responses.CreateOrder{Order_id: order_id, Products: fabrics}, nil
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
		status = ?
	WHERE order_id = ?`

	_, err := o.db.ExecContext(ctx, query, req.Status, req.Order_id)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderMySQL) UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	if err := utils.CheckOrderPayment(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	query := `
	UPDATE ORDERS 
	SET 
		is_payment = ?,
		status = ?
	WHERE order_id = ?`

	_, err := o.db.ExecContext(ctx, query, req.Is_payment, "waiting_assign", req.Order_id)
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

func (o *OrderMySQL) UpdatePrice(ctx context.Context, req *requests.UpdatePrice) error {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return err
	}

	query := `
	UPDATE ORDERS 
	SET 
		price = ?,
		status = ?
	WHERE order_id = ?
	`

	_, err := o.db.ExecContext(ctx, query, req.Price, "payment", req.Order_id)
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
		status = ?,
		tailor_id = ?,
		tailor_phone = ?,
		tailor_address = ?,
		due_date = ?
	WHERE order_id = ?
	`

	if _, err := o.db.ExecContext(ctx, query,
		"processing_user",
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

func (o *OrderMySQL) GetOrderByUserId(ctx context.Context, req *requests.UserID) ([]*responses.ShowOrder, error) {

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

func (o *OrderMySQL) CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.ProductProcess, error) {

	if err := utils.CheckOrderByID(o.db, ctx, req.Order_id); err != nil {
		return nil, err
	}

	// var res *responses.CheckProcess
	query := `
	SELECT 
		SUM(process_quantity) AS process_quantity, 
    	SUM(total_quantity) AS total_quantity
	FROM PRODUCTS WHERE created_by = ?
	`
	var product responses.ProductProcess
	err := o.db.GetContext(ctx, &product, query, req.Order_id)
	if err != nil {
		return nil, err
	}

	// res = &responses.CheckProcess{Is_ready: product.Process_quantity == product.Total_quantity}
	return &responses.ProductProcess{
		Is_ready:         product.Process_quantity == product.Total_quantity,
		Process_quantity: product.Process_quantity,
		Total_quantity:   product.Total_quantity,
	}, nil
}

func (o *OrderMySQL) CheckFabric(ctx context.Context, req *requests.Product, index string) (*responses.CheckFabric, error) {
	var res *responses.CheckFabric

	if err := utils.CheckFabricByID(o.db, ctx, req.Fabric_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		quantity 
	FROM FABRICS WHERE fabric_id = ?
	`

	var fabric responses.FabricQuantity
	err := o.db.GetContext(ctx, &fabric, query, req.Fabric_id)
	if err != nil {
		return nil, err
	}
	res = &responses.CheckFabric{Product_index: index, Quantity: fabric.Quantity-req.Total_quantity >= 0}

	return res, nil
}
