package controllers

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/requests"
	"github.com/SA-TailorStore/Kanok-API/responses"
	"github.com/SA-TailorStore/Kanok-API/usercases"
)

type userController struct {
	service usercases.UserUseCase
}

func NewUserController(service usercases.UserUseCase) usercases.UserUseCase {
	return &userController{
		service: service,
	}
}

// Login implements usercases.UserUseCase.
func (u *userController) Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error) {
	panic("unimplemented")
}

// Register implements usercases.UserUseCase.
func (u *userController) Register(ctx context.Context, req *requests.UserRegisterRequest) error {
	panic("unimplemented")
}
