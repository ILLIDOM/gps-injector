package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	// Setup Arango Connection
	arangoConfig := ArangoConfig{
		URL:      "https://arango.dev.network.garden", // golang http error when port is specified "http: server gave HTTP response to HTTPS client"
		User:     "root",
		Password: "jalapeno",
		Database: "jalapeno",
	}

	arango, err := NewArangoConnection(arangoConfig)
	if err != nil {
		log.Fatalf("Failed to create ArangoConnection: %v", err)
	}

	collectionExists, err := arango.db.CollectionExists(context.TODO(), "ls_node")
	if err != nil {
		log.Fatalf("Failed to check collection: %v", err)
	}

	fmt.Println(collectionExists)
}
