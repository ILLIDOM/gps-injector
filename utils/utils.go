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