package responses

type UserLoginResponse struct {
	User_id    string `json:"user_id"`
	Username   string `json:"username"`
	Token      string `json:"token"`
	Created_at string `json:"created_at"`
}

type UserResponse struct {
	User_id          string `json:"user_id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Display_name     string `json:"display_name"`
	User_profile_url string `json:"user_profile_url"`
	Role             string `json:"role"`
	Phone_number     string `json:"phone_number"`
	Address          string `json:"address"`
	Created_at       string `json:"created_at"`
}

type UsernameResponse struct {
	Username string `json:"username"`
}
