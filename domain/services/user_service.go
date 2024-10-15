package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v5"
)

type UserUseCase interface {
	GetAllUser(ctx context.Context) ([]*responses.Username, error)
	Login(ctx context.Context, req *requests.UserLogin) (*responses.UserJWT, error)
	Register(ctx context.Context, req *requests.UserRegister) error
	StoreRegister(ctx context.Context, req *requests.UserRegister) error
	GetByUsername(ctx context.Context, req *requests.Username) (*responses.Username, error)
	GetByJWT(ctx context.Context, req *requests.UserJWT) (*responses.User, error)
	GenerateToken(ctx context.Context, req *requests.UserJWT) (*responses.UserJWT, error)
	GetByID(ctx context.Context, req *requests.UserID) (*responses.User, error)
	UpdateAddress(ctx context.Context, req *requests.UserUpdate) error
	UploadImage(ctx context.Context, file interface{}, req *requests.UserUploadImage) (*responses.UserProUrl, error)
}

type userService struct {
	reposititory reposititories.UserRepository
	config       *configs.Config
	cloudinary   *cloudinary.Cloudinary
}

func NewUserService(reposititory reposititories.UserRepository, config *configs.Config, cld *cloudinary.Cloudinary) UserUseCase {
	return &userService{
		reposititory: reposititory,
		config:       config,
		cloudinary:   cld,
	}
}

func (u *userService) GetAllUser(ctx context.Context) ([]*responses.Username, error) {
	users, err := u.reposititory.GetAllUser(ctx)

	if err != nil {
		return nil, err
	}

	usernamesRes := make([]*responses.Username, 0)
	for _, user := range users {
		usernamesRes = append(usernamesRes, &responses.Username{
			Username: user.Username,
		})
	}

	return usernamesRes, err
}

func (u *userService) Register(ctx context.Context, req *requests.UserRegister) error {

	if err := utils.ValidateUsername(req.Username); err != nil {
		return err
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)
	return u.reposititory.CreateUser(ctx, req)
}

func (u *userService) StoreRegister(ctx context.Context, req *requests.UserRegister) error {

	if err := utils.ValidateUsername(req.Username); err != nil {
		return err
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)
	return u.reposititory.CreateTailor(ctx, req)
}

func (u *userService) Login(ctx context.Context, req *requests.UserLogin) (*responses.UserJWT, error) {

	res, err := u.reposititory.GetPasswordByUsername(ctx, &requests.Username{Username: req.Username})
	// Check if user exist
	if err != nil {
		return nil, err
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password)) != nil {
		return nil, exceptions.ErrWrongPassword
	}

	tokenString := utils.GenerateJWT(res.User_id)

	return &responses.UserJWT{
		Token: tokenString,
	}, nil
}

func (u *userService) GetByUsername(ctx context.Context, req *requests.Username) (*responses.Username, error) {

	err := u.reposititory.GetByUsername(ctx, req)

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

	user := &responses.Username{
		Username: req.Username,
	}

	return user, err
}

func (u *userService) GetByJWT(ctx context.Context, req *requests.UserJWT) (*responses.User, error) {
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

		return &responses.User{
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

func (u *userService) GenerateToken(ctx context.Context, req *requests.UserJWT) (*responses.UserJWT, error) {

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

	res, err := u.reposititory.GetUserByUserID(ctx, user_id)

	if err != nil {
		return nil, err
	}

	tokenString := utils.GenerateJWT(res.User_id)

	return &responses.UserJWT{
		Token: tokenString,
	}, err
}

func (u *userService) GetByID(ctx context.Context, req *requests.UserID) (*responses.User, error) {

	user, err := u.reposititory.GetUserByUserID(ctx, req)

	if err != nil {
		return nil, err
	}

	return &responses.User{
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

func (u *userService) UploadImage(ctx context.Context, file interface{}, req *requests.UserUploadImage) (*responses.UserProUrl, error) {

	user_id, err := utils.VerificationJWT(req.Token)
	if err != nil {
		return nil, err
	}

	res, _ := u.reposititory.GetUserByUserID(ctx, &requests.UserID{User_id: user_id})

	if res.User_profile_url != "-" {
		public_id, err := utils.ExtractPublicID(res.User_profile_url)
		if err != nil {
			return nil, err
		}
		_, err = u.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})

		if err != nil {
			return nil, err
		}

		err = u.reposititory.UpdateImage(ctx, &requests.UserUploadImage{
			Token: user_id,
			Image: "-",
		})
		if err != nil {
			return nil, err
		}

	}

	resCloud, err := u.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return nil, exceptions.ErrUploadImage
	}

	err = u.reposititory.UpdateImage(ctx, &requests.UserUploadImage{
		Token: user_id,
		Image: resCloud.SecureURL,
	})
	if err != nil {
		return nil, err
	}

	return &responses.UserProUrl{
		User_profile_url: resCloud.SecureURL,
		User_id:          user_id,
	}, err
}
