package entities

type Order struct {
	Order_id      string `json:"order_id" db:"order_id"`
	Status        string `json:"status" db:"status"`
	Is_payment    bool   `json:"is_payment" db:"is_payment"`
	Store_phone   string `json:"store_phone" db:"store_phone"`
	Store_address string `json:"store_address" db:"store_address"`
	User_phone    string `json:"user_phone" db:"user_phone"`
	User_address  string `json:"user_address" db:"user_address"`
	Due_date      string `json:"due_date" db:"due_date"`
	Create_by     string `json:"create_by" db:"create_by"`
	Created_at    string `json:"created_at" db:"created_at"`
}
