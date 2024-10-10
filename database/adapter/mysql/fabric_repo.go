package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/jmoiron/sqlx"
)

type FabricMySQL struct {
	db *sqlx.DB
}

func NewFabricMySQL(db *sqlx.DB) reposititories.FabricRepository {
	return &FabricMySQL{
		db: db,
	}
}

func (f *FabricMySQL) AddFabric(ctx context.Context, req *requests.AddFabric) error {
	query := `
	INSERT INTO FABRICS 
	(fabric_url,quantity) 
	VALUES (?,?)`

	_, err := f.db.QueryContext(ctx, query, req.Image, req.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (f *FabricMySQL) UpdateFabric(ctx context.Context, req *requests.UpdateFabric) error {
	query := `
	UPDATE FABRICS 
	SET 
		fabric_url = ?, 
		quantity = ? 
	WHERE fabric_id = ?`

	_, err := f.db.ExecContext(ctx, query, req.Image, req.Quantity, req.Fabric_id)
	if err != nil {
		return err
	}
	return nil
}

func (f *FabricMySQL) UpdateFabrics(ctx context.Context, req []*requests.UpdateFabrics) error {
	query := `
	UPDATE FABRICS 
	SET 
		quantity = ? 
	WHERE fabric_id = ?`

	for _, value := range req {

		_, err := f.db.ExecContext(ctx, query, value.Quantity, value.Fabric_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *FabricMySQL) DeleteFabric(ctx context.Context, req *requests.FabricID) error {

	query := `
	SELECT 
		fabric_id,
	FROM FABRICS WHERE fabric_id = ?
	`
	_, err := f.db.QueryContext(ctx, query, req.Fabric_id)
	if err != nil {
		return exceptions.ErrFabricNotFound
	}

	query = `DELETE FROM FABRICS WHERE fabric_id = ?`
	_, err = f.db.QueryContext(ctx, query, req.Fabric_id)
	if err != nil {
		return err
	}

	return nil
}

func (f *FabricMySQL) GetAllFabrics(ctx context.Context) ([]*responses.Fabric, error) {
	query := `
	SELECT 
		fabric_id, 
		fabric_url, 
		quantity 
	FROM FABRICS
	`

	rows, err := f.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fabrics []*responses.Fabric
	for rows.Next() {
		var fabric responses.Fabric
		err := rows.Scan(
			&fabric.Fabric_id,
			&fabric.Fabric_url,
			&fabric.Quantity)
		if err != nil {
			return nil, err
		}
		fabrics = append(fabrics, &fabric)
	}

	return fabrics, nil
}

func (f *FabricMySQL) GetFabricByID(ctx context.Context, req *requests.FabricID) (*responses.Fabric, error) {
	query := `
	SELECT 
		fabric_id, 
		fabric_url, 
		quantity 
	FROM FABRICS WHERE fabric_id = ?
	`

	var fabric responses.Fabric
	err := f.db.GetContext(ctx, &fabric, query, req.Fabric_id)
	if err != nil {
		return &fabric, err
	}

	return &fabric, nil
}
