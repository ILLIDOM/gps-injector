package arango

type LSNode struct {
	ID  string `json:"_id,omitempty"`
	Key string `json:"_key,omitempty"`
}

type LSNode_Coordinate struct {
	Key       string  `json:"_key,omitempty"`
	ID        string  `json:"_id,omitempty"`
	Latitude  float64 // `json:"latitude,omitempty"`
	Longitude float64 // `json:"longitude,omitempty"`
}
