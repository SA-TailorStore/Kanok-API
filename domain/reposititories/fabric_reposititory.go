package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
)

type FabricRepository interface {
	CreateFabric(ctx context.Context, req *requests.CreateFabric) error
}
