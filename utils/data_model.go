package utils

type LSNode struct {
	ID  string `json:"_id,omitempty"`
	Key string `json:"_key,omitempty"`
}

type LSNode_Coordinate struct {
	Key       string  `json:"_key,omitempty"`
	ID        string  `json:"_id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
