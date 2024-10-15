package entities

type Material struct {
	Material_id   string `json:"material_id" db:"material_id"`
	Material_name string `json:"material_name" db:"material_name"`
	Amount        int    `json:"amount" db:"amount"`
	Is_delete     int    `json:"is_delete" db:"is_delete"`
	Timestamp     string `json:"timestamp" db:"timestamp"`
}
