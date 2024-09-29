package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/entities"
	"github.com/SA-TailorStore/Kanok-API/requests"
	"github.com/SA-TailorStore/Kanok-API/responses"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	FindAllUser(ctx context.Context) ([]entities.User, error)
	FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error)
	GetUserByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UserResponse, error)
}
