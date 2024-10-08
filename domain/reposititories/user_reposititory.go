package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegister) error
	GetAllUser(ctx context.Context) ([]*responses.User, error)
	FindByUsername(ctx context.Context, req *requests.Username) error
	GetPasswordByUsername(ctx context.Context, req *requests.Username) (*responses.UserLogin, error)
	GetUserByUserID(ctx context.Context, req *requests.UserID) (*responses.User, error)
	UpdateAddress(ctx context.Context, req *requests.UserUpdate) error
	UploadImage(ctx context.Context, req *requests.UserUploadImage) error
}
