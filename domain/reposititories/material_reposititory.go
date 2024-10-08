package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
)

type MaterialRepository interface {
	CreateMaterial(ctx context.Context, req *requests.CreateMaterial) error
}
