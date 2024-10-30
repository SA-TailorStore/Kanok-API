package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
	"github.com/jmoiron/sqlx"
)

type MaterialMySQL struct {
	db *sqlx.DB
}

func NewMaterialMySQL(db *sqlx.DB) reposititories.MaterialRepository {
	return &MaterialMySQL{
		db: db,
	}
}

func (m *MaterialMySQL) AddMaterial(ctx context.Context, req *requests.AddMaterial) error {

	if err := utils.CheckNameDup(m.db, ctx, req.Material_name); err != nil {
		return err
	}

	query := `
	INSERT INTO MATERIALS 
	(material_name,amount)
	VALUES (?,?)`

	_, err := m.db.QueryContext(ctx, query, req.Material_name, req.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialMySQL) UpdateMaterial(ctx context.Context, req *requests.UpdateMaterial) error {

	if err := utils.CheckMaterialByID(m.db, ctx, req.Material_id); err != nil {
		return err
	}

	if err := utils.CheckNameDup(m.db, ctx, req.Material_name); err != nil {
		return err
	}

	query := `
	UPDATE MATERIALS 
	SET 
		material_name = ?, 
		amount = ? 
	WHERE material_id = ?`

	_, err := m.db.ExecContext(ctx, query, req.Material_name, req.Amount, req.Material_id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MaterialMySQL) DeleteMaterial(ctx context.Context, req *requests.MaterialID) error {

	if err := utils.CheckMaterialByID(m.db, ctx, req.Material_id); err != nil {
		return err
	}

	query := `DELETE FROM MATERIALS WHERE material_id = ?`
	_, err := m.db.QueryContext(ctx, query, req.Material_id)
	if err != nil {
		return err
	}

	return nil
}

func (m *MaterialMySQL) GetAllMaterials(ctx context.Context) ([]*responses.Material, error) {
	query := `
	SELECT 
		material_id, 
		material_name, 
		amount 
	FROM MATERIALS`

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []*responses.Material
	for rows.Next() {
		var material responses.Material
		err := rows.Scan(&material.Material_id, &material.Material_name, &material.Amount)
		if err != nil {
			return nil, err
		}
		materials = append(materials, &material)
	}
	return materials, nil
}

func (m *MaterialMySQL) GetMaterialByID(ctx context.Context, req *requests.MaterialID) (*responses.Material, error) {

	if err := utils.CheckMaterialByID(m.db, ctx, req.Material_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		material_id, 
		material_name, 
		amount 
	FROM MATERIALS WHERE material_id = ?`

	var material responses.Material
	err := m.db.GetContext(ctx, &material, query, req.Material_id)
	if err != nil {
		return &material, err
	}

	return &material, nil
}
