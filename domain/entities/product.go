package entities

type Product struct {
	Product_id       string `json:"product_id" db:"product_id"`
	Design_id        string `json:"design_id" db:"design_id"`
	Fabric_id        string `json:"fabric_id" db:"fabric_id"`
	Detail           string `json:"detail" db:"detail"`
	Size             string `json:"size" db:"size"`
	Process_quantity int    `json:"process_quantity" db:"process_quantity"`
	Max_quantity     int    `json:"max_quantity" db:"max_quantity"`
	Created_by       string `json:"created_by" db:"created_by"`
	Timestamp        string `json:"timestamp" db:"timestamp"`
}
