package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.OrderID, error)
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error
	UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error
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

func (o *orderService) CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.OrderID, error) {

	user_id, err := utils.VerificationJWT(req.Token)
	if err != nil {
		return nil, err
	}

	res, err := o.reposititory.CreateOrder(ctx, &requests.CreateOrder{Token: user_id})

	if err != nil {
		return res, err
	}

	return res, nil
}

func (o *orderService) UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error {

	err := o.reposititory.UpdateStatus(ctx, req)

	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error {

	err := o.reposititory.UpdatePayment(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
