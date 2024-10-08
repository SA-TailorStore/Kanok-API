package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
)

type FabricRepository interface {
	AddFabric(ctx context.Context, req *requests.AddFabric) error
	UpdateFabric(ctx context.Context, req *requests.UpdateFabric) error
	DeleteFabric(ctx context.Context, req *requests.FabricID) error
	GetFabricByID(ctx context.Context, req *requests.FabricID) error
	GetAllFabrics(ctx context.Context) error
}
