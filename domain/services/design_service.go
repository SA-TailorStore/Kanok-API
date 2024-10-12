package services

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
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

func (d *designService) AddDesign(ctx context.Context, file interface{}, req *requests.AddDesign) error {

	res, err := d.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return exceptions.ErrUploadImage
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

func (d *designService) UpdateDesign(ctx context.Context, file interface{}, req *requests.UpdateDesign) error {

	temp, err := d.reposititory.GetDesignByID(ctx, &requests.DesignID{Design_id: req.Design_id})
	if err != nil {
		return err
	}

	if req.Type == "" {
		req = &requests.UpdateDesign{
			Design_id: req.Design_id,
			Image:     temp.Design_url,
			Type:      temp.Type,
		}
	}

	if file != nil {
		if temp.Design_url != "-" {
			public_id, err := utils.ExtractPublicID(temp.Design_url)
			if err != nil {
				return err
			}
			_, err = d.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})

			if err != nil {
				return err
			}

			err = d.reposititory.UpdateDesign(ctx, &requests.UpdateDesign{
				Design_id: temp.Design_id,
				Image:     "-",
				Type:      temp.Type,
			})
			if err != nil {
				return err
			}
		}

		res, err := d.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
		if err != nil {
			return exceptions.ErrUploadImage
		}

		req = &requests.UpdateDesign{
			Design_id: req.Design_id,
			Image:     res.SecureURL,
			Type:      req.Type,
		}
	} else {
		req = &requests.UpdateDesign{
			Design_id: req.Design_id,
			Image:     temp.Design_url,
			Type:      req.Type,
		}
	}

	err = d.reposititory.UpdateDesign(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (d *designService) DeleteDesign(ctx context.Context, req *requests.DesignID) error {

	res, err := d.reposititory.GetDesignByID(ctx, &requests.DesignID{Design_id: req.Design_id})
	if err != nil {
		return exceptions.ErrDesignNotFound
	}

	public_id, err := utils.ExtractPublicID(res.Design_url)
	if err != nil {
		return err
	}

	_, err = d.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})
	if err != nil {
		return err
	}

	if err := d.reposititory.DeleteDesign(ctx, req); err != nil {
		return err
	}

	return nil
}

func (d *designService) GetAllDesigns(ctx context.Context) ([]*responses.Design, error) {

	designs, err := d.reposititory.GetAllDesigns(ctx)
	if err != nil {
		return nil, err
	}

	return designs, nil
}

func (d *designService) GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error) {

	design, err := d.reposititory.GetDesignByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return design, nil
}
