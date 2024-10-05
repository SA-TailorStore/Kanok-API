package entities

type Fabric struct {
	Fabric_id  string `json:"fabric_id" db:"fabric_id"`
	Fabric_url string `json:"fabric_url" db:"fabric_url"`
	Quantity   int    `json:"quantity" db:"quantity"`
	Create_at  string `json:"create_at" db:"create_at"`
}
