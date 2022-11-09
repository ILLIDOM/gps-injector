package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoConfig struct {
	URL      string
	User     string
	Password string
	Database string
}

type ArangoConn struct {
	db driver.Database
}

func NewArangoConnection(cfg ArangoConfig) (*ArangoConn, error) {
	// check configuration
	if cfg.URL == "" || cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
		return nil, errors.New("ArangoDB config has empty field")
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{cfg.URL},
	})
	if err != nil {
		return nil, errors.New("failed to create HTTP connection")
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(cfg.User, cfg.Password),
	})
	if err != nil {
		return nil, errors.New("failed to create ArangoDB client")
	}

	db, err := client.Database(context.TODO(), cfg.Database)
	if err != nil {
		return nil, errors.New("failed to open database")
	}

	return &ArangoConn{
		db: db,
	}, nil

}

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
