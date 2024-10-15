package entities

type Fabric struct {
	Fabric_id  string `json:"fabric_id" db:"fabric_id"`
	Fabric_url string `json:"fabric_url" db:"fabric_url"`
	Quantity   int    `json:"quantity" db:"quantity"`
	Is_delete  int    `json:"is_delete" db:"is_delete"`
	Timestamp  string `json:"timestamp" db:"timestamp"`
}
