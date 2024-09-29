package services

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/domain/usercases"
	"github.com/golang-jwt/jwt/v5"
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
	username := requests.UsernameRequest{
		Username: req.Username,
	}

	user, err := u.userRepo.GetUserByUsername(ctx, &username)

	// Check if user exist
	if err == exceptions.ErrUserNotFound {
		return nil, exceptions.ErrUserNotFound
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, exceptions.ErrLoginFailed
	}

	// Generate JWT token
	expireAt := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"user_id":  user.User_id,
		"username": user.Username,
		"exp":      expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(u.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &responses.UserLoginResponse{
		User_id:    user.User_id,
		Username:   user.Username,
		Token:      tokenString,
		Created_at: user.Created_at,
	}, nil
}

// Register implements usercases.UserUseCase.
func (u *userService) Register(ctx context.Context, req *requests.UserRegisterRequest) error {

	username := requests.UsernameRequest{
		Username: req.Username,
	}

	user, err := u.userRepo.FindByUsername(ctx, &username)

	if user == nil {
		return exceptions.ErrDuplicatedUsername
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

// FindByUsername implements usercases.UserUseCase.
func (u *userService) FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error) {
	user, err := u.userRepo.FindByUsername(ctx, req)

	if err == exceptions.ErrUserNotFound {
		return user, err
	}

	if user != nil {
		return user, exceptions.ErrDuplicatedUsername
	}

	return user, err
}
