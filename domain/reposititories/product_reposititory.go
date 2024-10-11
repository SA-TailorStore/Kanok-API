package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *requests.Product, order_id string) error
	GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.Product, error)
	GetProductByID(ctx context.Context, req *requests.ProductID) (*responses.Product, error)
}
