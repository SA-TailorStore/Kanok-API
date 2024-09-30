package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	GetAllUser(ctx context.Context) ([]entities.User, error)
	FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error)
	GetPasswordByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UserLoginResponse, error)
	GetUserByUserID(ctx context.Context, req *requests.UserID) (*responses.UserResponse, error)
}
