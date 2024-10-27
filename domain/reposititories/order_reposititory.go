package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.CreateOrder, error)
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error
	UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error
	UpdatePrice(ctx context.Context, req *requests.UpdatePrice) error
	UpdateTailor(ctx context.Context, req *requests.UpdateTailor) error
	UpdateTracking(ctx context.Context, req *requests.UpdateTracking) error
	GetOrderByUserId(ctx context.Context, req *requests.UserID) ([]*responses.ShowOrder, error)
	GetAllOrders(ctx context.Context) ([]*responses.ShowOrder, error)
	CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.ProductProcess, error)
	CheckFabric(ctx context.Context, req *requests.Product, index string) (*responses.CheckFabric, error)
}
