package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) error
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
}
