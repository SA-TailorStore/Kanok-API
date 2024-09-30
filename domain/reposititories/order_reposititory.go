package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) error
	GetOrderByID(ctx context.Context, req *requests.OrderIDRequest) error
}
