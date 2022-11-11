package utils

type LSNode struct {
	ID  string `json:"_id,omitempty"`
	Key string `json:"_key,omitempty"`
}

type LSNodeCoordinate struct {
	Key       string  `json:"_key,omitempty"`
	ID        string  `json:"_id,omitempty"`
	LsNodeKey string  `json:"ls_node_key"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
