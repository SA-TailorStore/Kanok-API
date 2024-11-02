package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type MaterialRepository interface {
	AddMaterial(ctx context.Context, req *requests.AddMaterial) error
	UpdateMaterial(ctx context.Context, req *requests.UpdateMaterial) error
	DeleteMaterial(ctx context.Context, req *requests.MaterialID) error
	GetAllMaterials(ctx context.Context) ([]*responses.Material, error)
	GetMaterialByID(ctx context.Context, req *requests.MaterialID) (*responses.Material, error)
}
