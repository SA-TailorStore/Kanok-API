package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/entities"
	"github.com/SA-TailorStore/Kanok-API/requests"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	FindAllUser(ctx context.Context) ([]entities.User, error)
	FindByUsername(ctx context.Context, email string) (*entities.User, error)
}
