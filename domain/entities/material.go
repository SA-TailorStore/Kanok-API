package entities

type Material struct {
	Material_id   string `json:"material_id" db:"material_id"`
	Material_name string `json:"material_name" db:"material_name"`
	Product_id    string `json:"product_id" db:"product_id"`
	Category      string `json:"category" db:"category"`
	Quantity      int    `json:"quantity" db:"quantity"`
	Timestamp     string `json:"timestamp" db:"timestamp"`
}
