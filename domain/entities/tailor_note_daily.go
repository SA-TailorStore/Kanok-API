package entities

type TailoringNoteDaily struct {
	Note_id   string `json:"note_id" db:"note_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Create_by string `json:"create_by" db:"create_by"`
	Create_at string `json:"create_at" db:"create_at"`
}
