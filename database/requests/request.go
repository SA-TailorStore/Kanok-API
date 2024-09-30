package requests

type UserRegisterRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Phone_number string `json:"phone_number" validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

type UserJWTRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserIDRequest struct {
	User_id string `json:"user_id" validate:"required"`
}

type OrderIDRequest struct {
	Order_id string `json:"order_id" validate:"required"`
}

type CreateOrderRequest struct {
	Store_phone   string `json:"store_phone" validate:"required"`
	Store_address string `json:"store_address" validate:"required"`
	User_phone    string `json:"user_phone" validate:"required"`
	User_address  string `json:"user_address" validate:"required"`
	Create_by     string `json:"create_by" validate:"required"`
}
