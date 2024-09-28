package entities

import "database/sql"

type User struct {
	ID               string         `json:"user_id" db:"user_id"`
	Username         string         `json:"username" db:"username"`
	Password         string         `json:"password" db:"password"`
	Display_name     sql.NullString `json:"display_name" db:"display_name"`
	User_profile_url sql.NullString `json:"user_profile_url" db:"user_profile_url"`
	Role             sql.NullString `json:"role" db:"role"`
	Phone_number     string         `json:"phone_number" db:"phone_number"`
	Address          sql.NullString `json:"address" db:"address"`
	Created_at       string         `json:"created_at" db:"created_at"`

	RoleString           string
	AddressString        string
	DisplayNameString    string
	UserProfileURLString string
}
