package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *requests.CreateProductRequest) error
	GetProductByOrderID(ctx context.Context, req *requests.OrderIDRequest) ([]*responses.ProductIDResponse, error)
}
