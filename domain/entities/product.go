package entities

type Product struct {
	Product_id   string `json:"product_id" db:"product_id"`
	Design_id    string `json:"design_id" db:"design_id"`
	Fabric_id    string `json:"fabric_id" db:"fabric_id"`
	Detail       string `json:"detail" db:"detail"`
	Size         string `json:"size" db:"size"`
	Quantity     int    `json:"quantity" db:"quantity"`
	Max_quantity int    `json:"max_quantity" db:"max_quantity"`
	Order_id     string `json:"order_id" db:"order_id"`
	Create_at    string `json:"create_at" db:"create_at"`
}
