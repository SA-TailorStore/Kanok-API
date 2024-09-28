package services

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/exceptions"
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

// FindAllUser implements usercases.UserUseCase.
func (u *userService) FindAllUser(ctx context.Context) ([]*responses.UsernameResponse, error) {
	users, err := u.userRepo.FindAllUser(ctx)

	if err != nil {
		return nil, err
	}

	usersResponse := make([]*responses.UsernameResponse, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, &responses.UsernameResponse{
			Username: user.Username,
		})
	}

	return usersResponse, err
}

// Login implements usercases.UserUseCase.
func (u *userService) Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserLoginResponse, error) {
	panic("unimplemented")
}

// Register implements usercases.UserUseCase.
func (u *userService) Register(ctx context.Context, req *requests.UserRegisterRequest) error {
	_, err := u.userRepo.FindByUsername(ctx, req.Username)

	if err == exceptions.ErrDuplicatedUsername {
		return err
	}

	if err == exceptions.ErrInvalidPassword {
		return err
	}

	if err == exceptions.ErrUsernameFormat {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)
	return u.userRepo.Create(ctx, req)

}
