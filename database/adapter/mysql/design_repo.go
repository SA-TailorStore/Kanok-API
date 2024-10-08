package mysql

import (
	"context"
	"database/sql"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
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

// AddDesign implements reposititories.DesignRepository.
func (d *DesignMySQL) AddDesign(ctx context.Context, req *requests.AddDesign) error {

	query := "INSERT INTO DESIGNS (design_url,type) VALUES (?,?)"

	_, err := d.db.QueryContext(ctx, query, req.Image, req.Type)
	if err != nil {
		return err
	}

	return nil
}

// UpdateDesign implements reposititories.DesignRepository.
func (d *DesignMySQL) UpdateDesign(ctx context.Context, req *requests.UpdateDesign) error {
	query := "UPDATE DESIGNS SET design_url = ?, type = ? WHERE design_id = ?"

	_, err := d.db.ExecContext(ctx, query, req.Image, req.Type, req.Design_ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDesign implements reposititories.DesignRepository.
func (d *DesignMySQL) DeleteDesign(ctx context.Context, req *requests.DesignID) error {
	query := "DELETE FROM DESIGNS WHERE design_id = ?"

	_, err := d.db.QueryContext(ctx, query, req.Design_id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllDesigns implements reposititories.DesignRepository.
func (d *DesignMySQL) GetAllDesigns(ctx context.Context) ([]*responses.Design, error) {
	rows, err := d.db.QueryContext(ctx, "SELECT design_id, design_url, type FROM DESIGNS")
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

// GetDesignByID implements reposititories.DesignRepository.
func (d *DesignMySQL) GetDesignByID(ctx context.Context, req *requests.DesignID) (*responses.Design, error) {
	var design responses.Design

	err := d.db.GetContext(ctx, &design, "SELECT design_id, design_url, type FROM DESIGNS WHERE design_id = ?", req.Design_id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, exceptions.ErrDesignNotFound
		default:
			return nil, err
		}

	}

	return &design, nil
}
