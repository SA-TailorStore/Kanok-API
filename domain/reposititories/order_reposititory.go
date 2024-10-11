package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type OrderRepository interface {
	CheckOrder(ctx context.Context, id string) error
	CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.OrderID, error)
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error
	UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error
}
