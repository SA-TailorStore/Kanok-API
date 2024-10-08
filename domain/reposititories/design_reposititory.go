package reposititories

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
)

type DesignRepository interface {
	CreateDesign(ctx context.Context, req *requests.CreateDesign) error
	GetDesignByID(ctx context.Context, req *requests.DesignID) error
}
