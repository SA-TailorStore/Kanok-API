package responses

type UserLoginResponse struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

type UserResponse struct {
	User_id          string `json:"user_id"`
	Username         string `json:"username"`
	Display_name     string `json:"display_name"`
	User_profile_url string `json:"user_profile_url"`
	Role             string `json:"role"`
	Phone_number     string `json:"phone_number"`
	Address          string `json:"address"`
	Timestamp        string `json:"timestamp"`
}

type UsernameResponse struct {
	Username string `json:"username"`
}

type UserJWT struct {
	Token string `json:"token"`
}

type OrderResponse struct {
	Order_id      string `json:"order_id"`
	Store_phone   string `json:"store_phone"`
	Store_address string `json:"store_address"`
	User_phone    string `json:"user_phone"`
	User_address  string `json:"user_address"`
	Created_by    string `json:"created_by"`
	Timestamp     string `json:"timestamp"`
}

type ProductResponse struct {
	Product_id       string `json:"product_id"`
	Design_id        string `json:"design_id"`
	Fabric_id        string `json:"fabric_id"`
	Detail           string `json:"detail"`
	Size             string `json:"size"`
	Process_quantity int    `json:"process_quantity"`
	Max_quantity     int    `json:"max_quantity"`
	Created_by       string `json:"created_by"`
	Timestamp        string `json:"timestamp"`
}

type ProductIDResponse struct {
	Product_id string `json:"product_id"`
}
