package services

import (
	"context"
	"strconv"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *requests.CreateProduct) error
	GetProductByID(ctx context.Context, req *requests.ProductID) (*responses.Product, error)
	GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.Product, error)
	UpdateProcessQuantity(ctx context.Context, req *requests.UpdateProcessQuantity) error
	GetAllProducts(ctx context.Context) ([]*responses.Product, error)
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

func (p *productService) CreateProduct(ctx context.Context, req *requests.CreateProduct) error {

	for index, value := range req.Products {
		err := p.reposititory.CreateProduct(ctx, &value, req.Order_id, strconv.Itoa(index+1))

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *productService) GetProductByOrderID(ctx context.Context, req *requests.OrderID) ([]*responses.Product, error) {
	products, err := p.reposititory.GetProductByOrderID(ctx, req)

	if err != nil {
		return nil, err
	}

	return products, err
}

func (p *productService) GetProductByID(ctx context.Context, req *requests.ProductID) (*responses.Product, error) {
	res, err := p.reposititory.GetProductByID(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *productService) GetAllProducts(ctx context.Context) ([]*responses.Product, error) {
	products, err := p.reposititory.GetAllProducts(ctx)

	if err != nil {
		return nil, err
	}

	return products, err
}
func (p *productService) UpdateProcessQuantity(ctx context.Context, req *requests.UpdateProcessQuantity) error {

	err := p.reposititory.UpdateProcessQuantity(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
