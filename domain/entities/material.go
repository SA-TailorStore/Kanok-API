package entities

type Material struct {
	Material_id   string `json:"material_id" db:"material_id"`
	Material_name string `json:"material_name" db:"material_name"`
	Category      string `json:"category" db:"category"`
	Quantity      int    `json:"quantity" db:"quantity"`
	Create_at     string `json:"create_at" db:"create_at"`
}
