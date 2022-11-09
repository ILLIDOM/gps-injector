package cmd

import (
	"context"
	"log"
	"os"

	"github.com/ILLIDOM/gps-injector/arango"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var input string // flag to specify input string

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
			collectionExists, err := arangoConn.Db.CollectionExists(context.TODO(), "ls_node_coordinates")
			if err != nil {
				log.Fatalf("Error checking collection: %v", err)
			}

			if collectionExists {
				log.Printf("Collection ls_node_coordinates already exists\nExiting\n")
				os.Exit(0)
			}

			// bytes, err = ioutil.ReadFile(input)
			// if err != nil {
			// 	log.Fatalf("Error reading input file: %v", err)
			// }

			// // allNodes :=

			// read from input file

			//insert into arangodb
		},
	}
	pushCmd.Flags().StringVar(&input, "i", "coordinates.json", "Flag to specify input file")
	pushCmd.MarkFlagRequired("i")
	return pushCmd
}
