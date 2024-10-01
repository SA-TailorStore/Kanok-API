package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *requests.CreateProductRequest) error
}
