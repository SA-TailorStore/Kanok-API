package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
)

type DesignRepository interface {
	AddDesign(ctx context.Context, req *requests.AddDesign) error
	UpdateDesign(ctx context.Context, req *requests.UpdateDesign) error
	DeleteDesign(ctx context.Context, req *requests.DesignID) error
	GetAllDesigns(ctx context.Context) ([]*responses.Design, error)
	GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error)
}
