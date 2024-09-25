package usercases

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/requests"
	"github.com/SA-TailorStore/Kanok-API/responses"
)

type UserUseCase interface {
	Register(ctx context.Context, req *requests.UserRegisterRequest) error
	Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error)
}
