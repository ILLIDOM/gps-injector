package utils

import (
	"encoding/json"
	"fmt"
)

func Print(any interface{}) {
	s, _ := json.MarshalIndent(any, "", "  ")
	fmt.Printf("%s\n\n", string(s))
}

func ToBytes(any interface{}) []byte {
	bytes, _ := json.MarshalIndent(any, "", "")
	return bytes
}

func ToLSNodeCoordinates(bytes []byte) []LSNode_Coordinate {
	var allNodes []LSNode_Coordinate
	_ = json.Unmarshal(bytes, &allNodes)
	return allNodes
}

func SetDefaultCoordinates(allNodes []LSNode_Coordinate) []LSNode_Coordinate {
	var nodes []LSNode_Coordinate

	for _, v := range allNodes {
		nodes = append(nodes,
			LSNode_Coordinate{
				Key:       v.Key,
				ID:        v.ID,
				Name:      v.Name,
				Latitude:  46.946296718731915,
				Longitude: 9.033762051602805,
			})
	}

	return nodes
}
