package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.OrderID, error)
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error
	UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error
	UpdateTailor(ctx context.Context, req *requests.UpdateTailor) error
	UpdateTracking(ctx context.Context, req *requests.UpdateTracking) error
	GetOrderByUserId(ctx context.Context, req *requests.UserID) ([]*responses.Order, error)
	GetAllOrders(ctx context.Context) ([]*responses.Order, error)
	CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.CheckProcess, error)
}
