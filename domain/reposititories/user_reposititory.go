package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	GetAllUser(ctx context.Context) ([]*responses.UserResponse, error)
	FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error)
	GetPasswordByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UserLoginResponse, error)
	GetUserByUserID(ctx context.Context, req *requests.UserIDRequest) (*responses.UserResponse, error)
}
