package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type DesignUseCase interface {
	AddDesign(ctx context.Context, file interface{}, req *requests.AddDesign) error
	UpdateDesign(ctx context.Context, file interface{}, req *requests.UpdateDesign) error
	DeleteDesign(ctx context.Context, req *requests.DesignID) error
	GetAllDesigns(ctx context.Context) ([]*responses.Design, error)
	GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error)
}

type designService struct {
	reposititory reposititories.DesignRepository
	cloudinary   *cloudinary.Cloudinary
	config       *configs.Config
}

func NewDesignService(reposititory reposititories.DesignRepository, config *configs.Config, cloudinary *cloudinary.Cloudinary) DesignUseCase {
	return &designService{
		reposititory: reposititory,
		cloudinary:   cloudinary,
		config:       config,
	}
}

// CreateDesign implements DesignUseCase.
func (d *designService) AddDesign(ctx context.Context, file interface{}, req *requests.AddDesign) error {

	res, err := d.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return err
	}

	req = &requests.AddDesign{
		Image: res.SecureURL,
		Type:  req.Type,
	}
	err = d.reposititory.AddDesign(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDesign implements DesignUseCase.
func (d *designService) UpdateDesign(ctx context.Context, file interface{}, req *requests.UpdateDesign) error {

	temp, err := d.reposititory.GetDesignByID(ctx, &requests.DesignID{Design_id: req.Design_ID})
	if err != nil {
		return err
	}
	if temp.Design_url != "-" {
		public_id, err := utils.ExtractPublicID(temp.Design_url)
		if err != nil {
			return err
		}
		_, err = d.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})

		if err != nil {
			return err
		}
		update := &requests.UpdateDesign{
			Design_ID: temp.Design_id,
			Image:     "-",
			Type:      temp.Type,
		}

		err = d.reposititory.UpdateDesign(ctx, update)
		if err != nil {
			return err
		}
	}

	res, err := d.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return err
	}

	req = &requests.UpdateDesign{
		Design_ID: req.Design_ID,
		Image:     res.SecureURL,
		Type:      req.Type,
	}

	err = d.reposititory.UpdateDesign(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDesign implements DesignUseCase.
func (d *designService) DeleteDesign(ctx context.Context, req *requests.DesignID) error {
	err := d.reposititory.DeleteDesign(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// GetAllDesigns implements DesignUseCase.
func (d *designService) GetAllDesigns(ctx context.Context) ([]*responses.Design, error) {
	var designs []*responses.Design
	return designs, nil
}

// GetDesignByID implements DesignUseCase.
func (d *designService) GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error) {
	design_id := &requests.DesignID{
		Design_id: req.Design_id,
	}

	design, err := d.reposititory.GetDesignByID(ctx, design_id)
	if err != nil {
		return nil, err
	}

	return design, nil
}
