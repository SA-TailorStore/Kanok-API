package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/jmoiron/sqlx"
)

type DesignMySQL struct {
	db *sqlx.DB
}

func NewDesignMySQL(db *sqlx.DB) reposititories.DesignRepository {
	return &DesignMySQL{
		db: db,
	}
}

func (d *DesignMySQL) AddDesign(ctx context.Context, req *requests.AddDesign) error {

	query := `
	INSERT INTO DESIGNS (design_url,type) 
	VALUES (?,?)`

	_, err := d.db.QueryContext(ctx, query, req.Image, req.Type)
	if err != nil {
		return err
	}

	return nil
}

func (d *DesignMySQL) UpdateDesign(ctx context.Context, req *requests.UpdateDesign) error {

	if err := utils.CheckDesignByID(d.db, ctx, req.Design_id); err != nil {
		return err
	}

	query := `
	UPDATE DESIGNS 
	SET 
		design_url = ?, 
		type = ? 
	WHERE design_id = ?`

	_, err := d.db.ExecContext(ctx, query, req.Image, req.Type, req.Design_id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DesignMySQL) DeleteDesign(ctx context.Context, req *requests.DesignID) error {

	if err := utils.CheckDesignByID(d.db, ctx, req.Design_id); err != nil {
		return err
	}

	query := `
	UPDATE DESIGNS 
	SET 
		is_delete = ? 
	WHERE design_id = ?`
	_, err := d.db.QueryContext(ctx, query, 1, req.Design_id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DesignMySQL) GetAllDesigns(ctx context.Context) ([]*responses.Design, error) {
	query := `
	SELECT
		design_id, 
		design_url, 
		type 
	FROM DESIGNS`

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var designs []*responses.Design
	for rows.Next() {
		var design responses.Design
		err := rows.Scan(&design.Design_id, &design.Design_url, &design.Type)
		if err != nil {
			return nil, err
		}
		designs = append(designs, &design)
	}

	return designs, nil
}

func (d *DesignMySQL) GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error) {

	if err := utils.CheckDesignByID(d.db, ctx, req.Design_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		design_id, 
		design_url, 
		type 
	FROM DESIGNS WHERE design_id = ?`

	var design responses.Design

	err := d.db.GetContext(ctx, &design, query, req.Design_id)

	if err != nil {
		return &design, err

	}

	return &design, nil
}
