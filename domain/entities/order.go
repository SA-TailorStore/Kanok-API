package entities

type Order struct {
	Order_id        string `json:"order_id" db:"order_id"`
	Status          string `json:"status" db:"status"`
	Is_payment      int    `json:"is_payment" db:"is_payment"`
	Store_phone     string `json:"store_phone" db:"store_phone"`
	Store_address   string `json:"store_address" db:"store_address"`
	Price           int    `json:"price" db:"price"`
	User_phone      string `json:"user_phone" db:"user_phone"`
	User_address    string `json:"user_address" db:"user_address"`
	Tracking_number string `json:"tracking_number" db:"tracking_number"`
	Tailor_id       string `json:"tailor_id" db:"tailor_id"`
	Due_date        string `json:"due_date" db:"due_date"`
	User_id         string `json:"user_id" db:"user_id"`
	Created_at      string `json:"created_at" db:"created_at"`
}
