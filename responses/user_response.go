package responses

type UserLoginResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
}
