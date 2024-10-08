package services

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/golang-jwt/jwt/v5"
)

type UserUseCase interface {
	GetAllUser(ctx context.Context) ([]*responses.UsernameResponse, error)
	Login(ctx context.Context, req *requests.UserLogin) (*responses.UserJWT, error)
	Register(ctx context.Context, req *requests.UserRegister) error
	FindByUsername(ctx context.Context, req *requests.Username) (*responses.UsernameResponse, error)
	FindByJWT(ctx context.Context, req *requests.UserJWT) (*responses.UserResponse, error)
	GenToken(ctx context.Context, req *requests.UserJWT) (*responses.UserJWT, error)
	FindByID(ctx context.Context, req *requests.UserID) (*responses.UserResponse, error)
	UpdateAddress(ctx context.Context, req *requests.UserUpdate) error
}

type userService struct {
	reposititory reposititories.UserRepository
	config       *configs.Config
}

func NewUserService(reposititory reposititories.UserRepository, config *configs.Config) UserUseCase {
	return &userService{
		reposititory: reposititory,
		config:       config,
	}
}

// FindAllUser implements usercases.UserUseCase.
func (u *userService) GetAllUser(ctx context.Context) ([]*responses.UsernameResponse, error) {
	users, err := u.reposititory.GetAllUser(ctx)

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
func (u *userService) Login(ctx context.Context, req *requests.UserLogin) (*responses.UserJWT, error) {
	username := &requests.Username{
		Username: req.Username,
	}

	user, err := u.reposititory.GetPasswordByUsername(ctx, username)
	// Check if user exist
	if err != nil {
		return nil, err
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, exceptions.ErrWrongPassword
	}

	// Generate JWT token
	expireAt := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"user_id": user.User_id,
		"exp":     expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(u.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &responses.UserJWT{
		Token: tokenString,
	}, nil
}

// Register implements usercases.UserUseCase.
func (u *userService) Register(ctx context.Context, req *requests.UserRegister) error {

	err := u.reposititory.FindByUsername(ctx, &requests.Username{Username: req.Username})
	if err != nil {
		return err
	}

	err = utils.ValidateUsername(req.Username)
	if err != nil {
		return err
	}

	err = utils.ValidatePassword(req.Password)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)
	return u.reposititory.Create(ctx, req)
}

// FindByUsername implements usercases.UserUseCase.
func (u *userService) FindByUsername(ctx context.Context, req *requests.Username) (*responses.UsernameResponse, error) {
	err := u.reposititory.FindByUsername(ctx, req)

	if err != nil {
		switch err {
		case exceptions.ErrUsernameDuplicated:
			return nil, err
		case exceptions.ErrUserNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	user := &responses.UsernameResponse{
		Username: req.Username,
	}

	return user, err
}

// FindByJWT implements UserUseCase.
func (u *userService) FindByJWT(ctx context.Context, req *requests.UserJWT) (*responses.UserResponse, error) {
	//JWT
	secret_key := []byte(u.config.JWTSecret)

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret_key, nil
	})

	// Check JWT
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		user_id := &requests.UserID{
			User_id: claims["user_id"].(string),
		}
		user, err := u.reposititory.GetUserByUserID(ctx, user_id)

		if err != nil {
			return nil, err
		}

		return &responses.UserResponse{
			User_id:          user.User_id,
			Username:         user.Username,
			Display_name:     user.Display_name,
			User_profile_url: user.User_profile_url,
			Role:             user.Role,
			Phone_number:     user.Phone_number,
			Address:          user.Address,
			Timestamp:        user.Timestamp,
		}, err
	} else {
		return nil, exceptions.ErrInvalidToken
	}
}

func (u *userService) GenToken(ctx context.Context, req *requests.UserJWT) (*responses.UserJWT, error) {
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

	user_id := &requests.UserID{
		User_id: id,
	}

	user, err := u.reposititory.GetUserByUserID(ctx, user_id)

	if err != nil {
		return nil, err
	}

	tokenString := utils.GenerateJWT(user.User_id)

	return &responses.UserJWT{
		Token: tokenString,
	}, err

}

func (u *userService) FindByID(ctx context.Context, req *requests.UserID) (*responses.UserResponse, error) {

	user, err := u.reposititory.GetUserByUserID(ctx, req)

	if err != nil {
		return nil, err
	}

	return &responses.UserResponse{
		User_id:          user.User_id,
		Username:         user.Username,
		Display_name:     user.Display_name,
		User_profile_url: user.User_profile_url,
		Role:             user.Role,
		Phone_number:     user.Phone_number,
		Address:          user.Address,
		Timestamp:        user.Timestamp,
	}, err
}

// UpdateAddress implements UserUseCase.
func (u *userService) UpdateAddress(ctx context.Context, req *requests.UserUpdate) error {
	user_id, err := utils.VerificationJWT(req.Token)

	if err != nil {
		switch err {
		case exceptions.ErrExpiredToken:
			return err
		case exceptions.ErrInvalidToken:
			return err
		default:
			return err
		}
	}

	req = &requests.UserUpdate{
		Token:        user_id,
		Display_name: req.Display_name,
		Phone_number: req.Phone_number,
		Address:      req.Address,
	}

	err = u.reposititory.UpdateAddress(ctx, req)

	if err != nil {
		switch err {
		case exceptions.ErrUserNotFound:
			return err
		default:
			return err
		}
	}

	return err
}
