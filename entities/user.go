package entities

type User struct {
	User_id          string `json:"user_id" db:"user_id"`
	Username         string `json:"username" db:"username"`
	Password         string `json:"password" db:"password"`
	Display_name     string `json:"display_name" db:"display_name"`
	User_profile_url string `json:"user_profile_url" db:"user_profile_url"`
	Role             string `json:"role" db:"role"`
	Phone_number     string `json:"phone_number" db:"phone_number"`
	Address          string `json:"address" db:"address"`
	Created_at       string `json:"created_at" db:"created_at"`
}
