package arango

import (
	"context"
	"errors"
	"log"

	"github.com/ILLIDOM/gps-injector/utils"
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
	Db driver.Database
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
		Db: db,
	}, nil

}

func DeleteCollection(arangoConn *ArangoConn, name string) error {
	coll, err := arangoConn.Db.Collection(context.TODO(), name)
	if err != nil {
		return err
	}
	err = coll.Remove(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func CreateCollection(arangoConn *ArangoConn, name string) driver.Collection {
	options := driver.CreateCollectionOptions{}
	col, err := arangoConn.Db.CreateCollection(context.TODO(), name, &options)
	if err != nil {
		log.Fatalf("Error creating collection %s: %v", name, err)
	}
	return col
}

func GetAllLSNodes(ctx context.Context, arangoConn *ArangoConn) []utils.LSNodeCoordinate {
	queryString := "FOR d in ls_node RETURN d"

	cursor := queryArango(ctx, arangoConn, queryString)
	var documents []utils.LSNodeCoordinate
	for {
		var document utils.LSNodeCoordinate
		doc, err := cursor.ReadDocument(ctx, &document)
		if !validDocument(doc, err) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}

func queryArango(ctx context.Context, arangoConn *ArangoConn, queryString string) driver.Cursor {
	cursor, err := arangoConn.Db.Query(ctx, queryString, nil)
	if err != nil {
		log.Fatalf("Error querying arangodb: %v", err)
	}
	defer cursor.Close()
	return cursor
}

func validDocument(document driver.DocumentMeta, err error) bool {
	if driver.IsNoMoreDocuments(err) {
		return false
	}
	if err != nil {
		log.Fatalf("Failed to read from ArangoDb: %v", err)
	}
	return true
}
