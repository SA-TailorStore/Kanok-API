package entities

type Design struct {
	Design_id  string `json:"design_id" db:"design_id"`
	Type       string `json:"type" db:"type"`
	Design_url string `json:"design_url" db:"design_url"`
	Created_at string `json:"created_at" db:"created_at"`
}
