package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.OrderID, error)
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error
	UpdatePayment(ctx context.Context, req *requests.UpdatePayment) error
	GetOrderByJWT(ctx context.Context, req *requests.UserJWT) ([]*responses.Order, error)
	GetAllOrders(ctx context.Context) ([]*responses.Order, error)
	StoreAssign(ctx context.Context, req *requests.StoreAssign) error
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

func (o *orderService) GetOrderByJWT(ctx context.Context, req *requests.UserJWT) ([]*responses.Order, error) {
	id, err := utils.VerificationJWT(req.Token)

	if err != nil {
		switch err {
		case exceptions.ErrExpiredToken:
			return nil, err
		case exceptions.ErrInvalidToken:
			return nil, err
		default:
			return nil, err
		}
	}

	res, err := o.reposititory.GetOrderByUserId(ctx, &requests.UserID{User_id: id})

	if err != nil {
		return res, err
	}

	return res, nil
}

func (o *orderService) GetAllOrders(ctx context.Context) ([]*responses.Order, error) {

	res, err := o.reposititory.GetAllOrders(ctx)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (o *orderService) StoreAssign(ctx context.Context, req *requests.StoreAssign) error {

	err := o.reposititory.StoreAssign(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
