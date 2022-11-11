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

func ToLSNodeCoordinates(bytes []byte) []LSNodeCoordinate {
	var allNodes []LSNodeCoordinate
	_ = json.Unmarshal(bytes, &allNodes)
	return allNodes
}

func SetDefaultCoordinates(allNodes []LSNodeCoordinate) []LSNodeCoordinate {
	var nodes []LSNodeCoordinate

	for _, v := range allNodes {
		nodes = append(nodes,
			LSNodeCoordinate{
				Key:       v.Key,
				ID:        v.ID,
				LsNodeKey: v.Key,
				Latitude:  46.946296718731915,
				Longitude: 9.033762051602805,
			})
	}

	return nodes
}
