package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *requests.UserRegister) error
	CreateTailor(ctx context.Context, req *requests.UserRegister) error
	GetAllUsers(ctx context.Context, req *requests.UserRole) ([]*responses.User, error)
	GetAllTailors(ctx context.Context, req *requests.UserRole) ([]*responses.UserTailor, error)
	GetByUsername(ctx context.Context, req *requests.Username) error
	GetPasswordByUsername(ctx context.Context, req *requests.Username) (*responses.UserLogin, error)
	GetUserByUserID(ctx context.Context, req *requests.UserID) (*responses.User, error)
	UpdateAddress(ctx context.Context, req *requests.UserUpdate) error
	UpdateImage(ctx context.Context, req *requests.UserUploadImage) error
}
