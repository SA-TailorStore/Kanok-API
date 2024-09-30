package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) error
	GetOrderByID(ctx context.Context, req *requests.OrderIDRequest) error
}

type orderService struct {
	reposititory reposititories.OrderRepository
	config       *configs.Config
}

func NewOrderService(reposititory reposititories.OrderRepository, config *configs.Config) OrderUseCase {
	return &orderService{
		reposititory: reposititory,
		config:       config,
	}
}

// GetOrder implements OrderUseCase.
func (o *orderService) GetOrderByID(ctx context.Context, req *requests.OrderIDRequest) error {

	err := o.reposititory.GetOrderByID(ctx, req)

	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return exceptions.ErrOrderNotFound
		default:
			return err
		}

	}

	return err
}

// CreateOrder implements OrderUseCase.
func (o *orderService) CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) error {

	err := o.reposititory.CreateOrder(ctx, req)

	if err != nil {
		switch err {
		case exceptions.ErrOrderNotFound:
			return exceptions.ErrOrderNotFound
		default:
			return err
		}

	}

	return err
}
