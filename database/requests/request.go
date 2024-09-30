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

type UserJWT struct {
	Token string `json:"token" validate:"required"`
}

type UserID struct {
	User_id string `json:"user_id" validate:"required"`
}
