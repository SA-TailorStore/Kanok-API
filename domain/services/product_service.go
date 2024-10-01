package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *requests.CreateProductRequest) error
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
func (p *productService) CreateProduct(ctx context.Context, req *requests.CreateProductRequest) error {
	err := p.reposititory.CreateProduct(ctx, req)

	if err != nil {
		return err
	}

	return err
}
