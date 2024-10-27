package services

import (
	"context"
	"mime/multipart"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.CreateOrder, error)
	GetOrderByID(ctx context.Context, req *requests.OrderID) (*responses.Order, error)
	UpdateStatus(ctx context.Context, req *requests.UpdateStatus) error
	UpdatePayment(ctx context.Context, req *requests.UpdatePayment, file multipart.File) error
	UpdateTailor(ctx context.Context, req *requests.UpdateTailor) error
	UpdateTracking(ctx context.Context, req *requests.UpdateTracking) error
	GetOrderByJWT(ctx context.Context, req *requests.UserJWT) ([]*responses.ShowOrder, error)
	GetAllOrders(ctx context.Context) ([]*responses.ShowOrder, error)
	CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.CheckProcess, error)
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

func (o *orderService) CreateOrder(ctx context.Context, req *requests.CreateOrder) (*responses.CreateOrder, error) {

	user_id, err := utils.VerificationJWT(req.Token)
	if err != nil {
		return nil, err
	}

	res, err := o.reposititory.CreateOrder(ctx, &requests.CreateOrder{Token: user_id, Products: req.Products})

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

func (o *orderService) UpdatePayment(ctx context.Context, req *requests.UpdatePayment, file multipart.File) error {
	// Decode รูปภาพ
	img, err := utils.DecodeImage(file)
	if err != nil {
		return err
	}

	// อ่าน QR code
	codes, err := utils.ReadQRCode(img)
	if err != nil {
		return err
	}

	if err := utils.ValidateSlip(codes); err != nil {
		return err
	}
	req = &requests.UpdatePayment{Order_id: req.Order_id, Is_payment: 1}
	if err := o.reposititory.UpdatePayment(ctx, req); err != nil {
		return err
	}

	return nil
}

func (o *orderService) UpdateTracking(ctx context.Context, req *requests.UpdateTracking) error {

	err := o.reposititory.UpdateTracking(ctx, req)

	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) GetOrderByJWT(ctx context.Context, req *requests.UserJWT) ([]*responses.ShowOrder, error) {
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

func (o *orderService) GetAllOrders(ctx context.Context) ([]*responses.ShowOrder, error) {

	res, err := o.reposititory.GetAllOrders(ctx)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (o *orderService) UpdateTailor(ctx context.Context, req *requests.UpdateTailor) error {

	err := o.reposititory.UpdateTailor(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.CheckProcess, error) {

	res, err := o.reposititory.CheckProcess(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
