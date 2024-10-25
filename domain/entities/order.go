package entities

type Order struct {
	Order_id        string `json:"order_id" db:"order_id"`
	Status          string `json:"status" db:"status"`
	Is_payment      int    `json:"is_payment" db:"is_payment"`
	Store_phone     string `json:"store_phone" db:"store_phone"`
	Store_address   string `json:"store_address" db:"store_address"`
	Price           int    `json:"price" db:"price"`
	Tailor_id       string `json:"tailor_id" db:"tailor_id"`
	Tailor_phone    string `json:"tailor_phone" db:"tailor_phone"`
	Tailor_address  string `json:"tailor_address" db:"tailor_address"`
	Tracking_number string `json:"tracking_number" db:"tracking_number"`
	Due_date        string `json:"due_date" db:"due_date"`
	User_phone      string `json:"user_phone" db:"user_phone"`
	User_address    string `json:"user_address" db:"user_address"`
	Create_by       string `json:"create_by" db:"create_by"`
	Timestamp       string `json:"timestamp" db:"timestamp"`
}
