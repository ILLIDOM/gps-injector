package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ILLIDOM/gps-injector/arango"
	"github.com/ILLIDOM/gps-injector/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	COORDINATES_COLLECTION string = "ls_node_coordinates"
)

var input string   // flag to specify input string
var overwrite bool // flag to specify if collection shall be overwritten if it already exists

func NewPushCmd() *cobra.Command {
	pushCmd := &cobra.Command{
		Use:     "push",
		Short:   "push and save nodes with coordindates from input file into arango ls_node_coordinates collection",
		Example: "gps-injector push -i coordinates.json",

		Run: func(cmd *cobra.Command, args []string) {
			// Create Arango Configuration
			arangoConfig := arango.ArangoConfig{
				URL:      viper.GetString("ARANGO_HOST"), // golang http error when port is specified "http: server gave HTTP response to HTTPS client"
				User:     viper.GetString("ARANGO_USER"),
				Password: viper.GetString("ARANGO_PASSWORD"),
				Database: viper.GetString("ARANGO_DB"),
			}

			// create Arango Connection
			arangoConn, err := arango.NewArangoConnection(arangoConfig)
			if err != nil {
				log.Fatalf("Failed to create ArangoConnection: %v", err)
			}

			// check if ls_node_coordinates collection exists
			collectionExists, err := arangoConn.Db.CollectionExists(context.TODO(), COORDINATES_COLLECTION)
			// arangodb returns 404 if collection doesnt exist
			if err != nil && !strings.Contains(err.Error(), "Unsupported content type 'text/plain; charset=utf-8'") {
				log.Fatalf("Failed to check collection: %v", err)
			}

			if collectionExists && !overwrite {
				log.Printf("Collection %s already exists. --overwrite is false\nExiting\n", COORDINATES_COLLECTION)
				os.Exit(0)
			}

			if collectionExists && overwrite {
				// overwrite collection
				err := arango.DeleteCollection(arangoConn, COORDINATES_COLLECTION)
				if err != nil {
					log.Fatalf("Error removing collection: %v", err)
				}
			}

			// create collection
			col := arango.CreateCollection(arangoConn, COORDINATES_COLLECTION)

			// read input file
			bytes, err := ioutil.ReadFile(input)
			if err != nil {
				log.Fatalf("Error reading input file: %v", err)
			}

			allNodes := utils.ToLSNodeCoordinates(bytes)

			// create documents in ls_node_coordiates collection
			_, errs, err := col.CreateDocuments(context.TODO(), allNodes)
			// errs contains possible errors for each documents trying to create
			if errs[0] != nil {
				log.Fatalf("Error writing document errors: %v", errs)
			}
			if err != nil {
				log.Fatalf("Error writing documents: %v", err)
			}

			fmt.Printf("Successfully added %s collection!", COORDINATES_COLLECTION)
		},
	}
	pushCmd.Flags().StringVarP(&input, "input", "i", "", "Flag to specify input file")
	pushCmd.MarkFlagRequired("input")
	pushCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "Flag to specify to overwrite collection")
	return pushCmd
}
