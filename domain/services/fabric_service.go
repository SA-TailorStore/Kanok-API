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

type FabricUseCase interface {
	AddFabric(ctx context.Context, file interface{}, req *requests.AddFabric) error
	UpdateFabric(ctx context.Context, file interface{}, req *requests.UpdateFabric) error
	UpdateFabrics(ctx context.Context, req []*requests.UpdateFabrics) error
	DeleteFabric(ctx context.Context, req *requests.FabricID) error
	GetFabricByID(ctx context.Context, req *requests.FabricID) (*responses.Fabric, error)
	GetAllFabrics(ctx context.Context) ([]*responses.Fabric, error)
}

type fabricService struct {
	reposititory reposititories.FabricRepository
	cloudinary   *cloudinary.Cloudinary
	config       *configs.Config
}

func NewFabricService(reposititory reposititories.FabricRepository, config *configs.Config, cloudinary *cloudinary.Cloudinary) FabricUseCase {
	return &fabricService{
		reposititory: reposititory,
		cloudinary:   cloudinary,
		config:       config,
	}
}

func (f *fabricService) AddFabric(ctx context.Context, file interface{}, req *requests.AddFabric) error {

	res, err := f.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return exceptions.ErrUploadImage
	}

	err = f.reposititory.AddFabric(ctx, &requests.AddFabric{
		Image:    res.SecureURL,
		Quantity: req.Quantity,
	})
	if err != nil {
		return err
	}

	return nil
}

func (f *fabricService) UpdateFabric(ctx context.Context, file interface{}, req *requests.UpdateFabric) error {

	temp, err := f.reposititory.GetFabricByID(ctx, &requests.FabricID{Fabric_id: req.Fabric_id})
	if err != nil {
		return err
	}

	if req.Quantity == 0 {
		req = &requests.UpdateFabric{
			Fabric_id: req.Fabric_id,
			Image:     temp.Fabric_url,
			Quantity:  temp.Quantity,
		}
	}

	if file != nil {
		if temp.Fabric_url != "-" {
			public_id, err := utils.ExtractPublicID(temp.Fabric_url)
			if err != nil {
				return err
			}
			_, err = f.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})

			if err != nil {
				return err
			}

			err = f.reposititory.UpdateFabric(ctx, &requests.UpdateFabric{
				Fabric_id: temp.Fabric_id,
				Image:     "-",
				Quantity:  temp.Quantity,
			})
			if err != nil {
				return err
			}
		}

		res, err := f.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
		if err != nil {
			return exceptions.ErrUploadImage
		}

		req = &requests.UpdateFabric{
			Fabric_id: req.Fabric_id,
			Image:     res.SecureURL,
			Quantity:  req.Quantity,
		}
	} else {
		req = &requests.UpdateFabric{
			Fabric_id: req.Fabric_id,
			Image:     temp.Fabric_url,
			Quantity:  req.Quantity,
		}
	}

	err = f.reposititory.UpdateFabric(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (f *fabricService) UpdateFabrics(ctx context.Context, req []*requests.UpdateFabrics) error {
	err := f.reposititory.UpdateFabrics(ctx, req)

	if err != nil {
		return err
	}

	return nil
}

func (f *fabricService) DeleteFabric(ctx context.Context, req *requests.FabricID) error {

	res, err := f.reposititory.GetFabricByID(ctx, &requests.FabricID{Fabric_id: req.Fabric_id})
	if err != nil {
		return exceptions.ErrFabricNotFound
	}

	public_id, err := utils.ExtractPublicID(res.Fabric_url)
	if err != nil {
		return err
	}

	_, err = f.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: public_id})
	if err != nil {
		return err
	}

	if err := f.reposititory.DeleteFabric(ctx, req); err != nil {
		return err
	}

	return nil
}

func (f *fabricService) GetAllFabrics(ctx context.Context) ([]*responses.Fabric, error) {
	res, err := f.reposititory.GetAllFabrics(ctx)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (f *fabricService) GetFabricByID(ctx context.Context, req *requests.FabricID) (*responses.Fabric, error) {
	res, err := f.reposititory.GetFabricByID(ctx, req)

	if err != nil {
		return res, err
	}

	return res, nil
}
