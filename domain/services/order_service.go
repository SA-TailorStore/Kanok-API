package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

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
	UpdatePrice(ctx context.Context, req *requests.UpdatePrice) error
	GetOrderByJWT(ctx context.Context, req *requests.UserJWT) ([]*responses.ShowOrder, error)
	GetAllOrders(ctx context.Context) ([]*responses.ShowOrder, error)
	CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.ProductProcess, error)
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
	s := utils.GetStringQR(codes)
	resp, _ := utils.SendString(s)

	cur_order, err := o.reposititory.GetOrderByID(ctx, &requests.OrderID{Order_id: req.Order_id})
	if err != nil {
		return err
	}

	if data, ok := resp["data"].(map[string]interface{}); ok {
		if code, ok := resp["code"].(float64); ok {
			fmt.Println(code)
			switch code {
			case 1012:
				return exceptions.ErrSlipIsDup
			case 1013:
				return exceptions.ErrWrongAmount
			case 1014:
				return exceptions.ErrWrongAccount
			default:
				return exceptions.ErrWrongSlip
			}
		} else {
			if amount, ok := data["amount"].(float64); ok && amount == cur_order.Price {
				req = &requests.UpdatePayment{Order_id: req.Order_id, Is_payment: 1}
			} else {
				return exceptions.ErrAmountIsWrong
			}
		}

	}

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

func (o *orderService) UpdatePrice(ctx context.Context, req *requests.UpdatePrice) error {

	if req.Price <= 0 {
		return exceptions.ErrPriceIsValid
	}

	err := o.reposititory.UpdatePrice(ctx, req)

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

	layout := "2006-01-02T15:04:05.000Z"
	parsedDate, err := time.Parse(layout, req.Due_date)

	if err != nil {
		return exceptions.ErrDateInvalid
	}
	duration := time.Since(parsedDate) * -1

	if duration <= 72*time.Hour {
		return exceptions.ErrDateToLow
	}

	req = &requests.UpdateTailor{Order_id: req.Order_id, Tailor_id: req.Tailor_id, ParseDate: parsedDate}
	if err := o.reposititory.UpdateTailor(ctx, req); err != nil {
		return err
	}

	return nil
}

func (o *orderService) CheckProcess(ctx context.Context, req *requests.OrderID) (*responses.ProductProcess, error) {

	res, err := o.reposititory.CheckProcess(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
