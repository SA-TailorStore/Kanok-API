package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) error
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
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

func (o *orderService) GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error) {

	res, err := o.reposititory.GetOrderByID(ctx, req)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (o *orderService) CreateOrder(ctx context.Context, req *requests.CreateOrder) error {

	err := o.reposititory.CreateOrder(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
