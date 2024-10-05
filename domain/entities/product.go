package entities

type Product struct {
	Product_id string `json:"product_id" db:"product_id"`
	Detail     string `json:"detail" db:"detail"`
	Size       string `json:"size" db:"size"`
	Quantity   int    `json:"quantity" db:"quantity"`
	Create_by  string `json:"create_by" db:"create_by"`
	Create_at  string `json:"create_at" db:"create_at"`
}
