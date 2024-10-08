package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *requests.CreateProduct) error
	GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.ProductIDResponse, error)
}

type productService struct {
	reposititory reposititories.ProductRepository
	config       *configs.Config
}

func NewProductService(reposititory reposititories.ProductRepository, config *configs.Config) ProductUsecase {
	return &productService{
		reposititory: reposititory,
		config:       config,
	}
}

// CreateProduct implements ProductUsecase.
func (p *productService) CreateProduct(ctx context.Context, req *requests.CreateProduct) error {
	err := p.reposititory.CreateProduct(ctx, req)

	if err != nil {
		switch err {
		case exceptions.ErrProductNotFound:
			return exceptions.ErrProductNotFound
		default:
			return err
		}
	}

	return err
}

// GetProductByOrderID implements ProductUsecase.
func (p *productService) GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.ProductIDResponse, error) {
	products, err := p.reposititory.GetProductByOrderID(ctx, req)

	if err != nil {
		return nil, err
	}

	return products, err
}
