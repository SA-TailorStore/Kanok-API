package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/reposititories"
	"github.com/SA-TailorStore/Kanok-API/requests"
	"github.com/SA-TailorStore/Kanok-API/responses"
	"github.com/SA-TailorStore/Kanok-API/usercases"
)

type userService struct {
	userRepo reposititories.UserRepository
	config   *configs.Config
}

func NewUserService(userRepo reposititories.UserRepository, config *configs.Config) usercases.UserUseCase {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}

// Login implements usercases.UserUseCase.
func (u *userService) Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error) {
	panic("unimplemented")
}

// Register implements usercases.UserUseCase.
func (u *userService) Register(ctx context.Context, req *requests.UserRegisterRequest) error {
	panic("unimplemented")
}

func (u *userService) GetUsers(ctx context.Context) ([]responses.UserResponse, error) {
	panic("unimplemented")
}
