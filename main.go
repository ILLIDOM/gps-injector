package main

import (
	"fmt"
	"os"

	"github.com/ILLIDOM/gps-injector/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// // Setup Arango Connection
	// arangoConfig := arango.ArangoConfig{
	// 	URL:      "https://arango.dev.network.garden", // golang http error when port is specified "http: server gave HTTP response to HTTPS client"
	// 	User:     "root",
	// 	Password: "jalapeno",
	// 	Database: "jalapeno",
	// }

	// arangoConn, err := arango.NewArangoConnection(arangoConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to create ArangoConnection: %v", err)
	// }

	// collectionExists, err := arangoConn.db.CollectionExists(context.TODO(), "ls_node")
	// if err != nil {
	// 	log.Fatalf("Failed to check collection: %v", err)
	// }

	// fmt.Println(collectionExists)
}
