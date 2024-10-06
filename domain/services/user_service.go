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
	"github.com/golang-jwt/jwt/v5"
)

type UserUseCase interface {
	Register(ctx context.Context, req *requests.UserRegisterRequest) error
	Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserJWT, error)
	GetAllUser(ctx context.Context) ([]*responses.UsernameResponse, error)
	FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error)
	FindByJWT(ctx context.Context, req *requests.UserJWTRequest) (*responses.UserResponse, error)
	GenToken(ctx context.Context, req *requests.UserJWTRequest) (*responses.UserJWT, error)
	FindByID(ctx context.Context, req *requests.UserIDRequest) (*responses.UserResponse, error)
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
func (u *userService) Login(ctx context.Context, req *requests.UserLoginRequest) (*responses.UserJWT, error) {
	username := &requests.UsernameRequest{
		Username: req.Username,
	}

	user, err := u.reposititory.GetPasswordByUsername(ctx, username)
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
func (u *userService) Register(ctx context.Context, req *requests.UserRegisterRequest) error {

	username := requests.UsernameRequest{
		Username: req.Username,
	}

	user, err := u.reposititory.FindByUsername(ctx, &username)

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
	return u.reposititory.Create(ctx, req)

}

// FindByUsername implements usercases.UserUseCase.
func (u *userService) FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error) {
	user, err := u.reposititory.FindByUsername(ctx, req)

	if err == exceptions.ErrUserNotFound {
		return user, err
	}

	if user != nil {
		return user, exceptions.ErrDuplicatedUsername
	}

	return user, err
}

// FindByJWT implements UserUseCase.
func (u *userService) FindByJWT(ctx context.Context, req *requests.UserJWTRequest) (*responses.UserResponse, error) {
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
		user_id := &requests.UserIDRequest{
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
			Created_at:       user.Created_at,
		}, err
	} else {
		return nil, exceptions.ErrInvalidToken
	}
}

func (u *userService) GenToken(ctx context.Context, req *requests.UserJWTRequest) (*responses.UserJWT, error) {
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
		user_id := &requests.UserIDRequest{
			User_id: claims["user_id"].(string),
		}
		user, err := u.reposititory.GetUserByUserID(ctx, user_id)

		if err != nil {
			return nil, err
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
		}, err
	} else {
		return nil, exceptions.ErrInvalidToken
	}
}

func (u *userService) FindByID(ctx context.Context, req *requests.UserIDRequest) (*responses.UserResponse, error) {

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
		Created_at:       user.Created_at,
	}, err

}
