package entities

type Design struct {
	Design_id  string `json:"design_id" db:"design_id"`
	Design_url string `json:"design_url" db:"design_url"`
	Type       string `json:"type" db:"type"`
	Timestamp  string `json:"timestamp" db:"timestamp"`
}
