package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type FabricRepository interface {
	AddFabric(ctx context.Context, req *requests.AddFabric) error
	UpdateFabric(ctx context.Context, req *requests.UpdateFabric) error
	UpdateFabrics(ctx context.Context, req []*requests.UpdateFabrics) error
	DeleteFabric(ctx context.Context, req *requests.FabricID) error
	GetFabricByID(ctx context.Context, req *requests.FabricID) (*responses.Fabric, error)
	GetAllFabrics(ctx context.Context) ([]*responses.Fabric, error)
}
