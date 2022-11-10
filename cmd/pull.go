package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ILLIDOM/gps-injector/arango"
	"github.com/ILLIDOM/gps-injector/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	NODE_COLLECTION string = "ls_node"
)

var output string // flag to specify output file

func NewPullCmd() *cobra.Command {
	pullCmd := &cobra.Command{
		Use:     "pull",
		Short:   "get and save yaml for coordinates",
		Example: "gps-injector pull -o coordinates.json",

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

			// check if ls_node collection exists
			collectionExists, err := arangoConn.Db.CollectionExists(context.TODO(), NODE_COLLECTION)
			if err != nil {
				log.Fatalf("Failed to check collection: %v", err)
			}

			if !collectionExists {
				log.Fatalf("%s collection does not exist. is Jalapeno up and running?", NODE_COLLECTION)
			}

			// fetch all nodes from arangodb
			allNodes := arango.GetAllLSNodes(context.TODO(), arangoConn)

			// write all nodes into output file
			err = ioutil.WriteFile(output, utils.ToBytes(allNodes), 0644)
			if err != nil {
				log.Fatalf("Error writing output file: %v", err)
			}

			fmt.Printf("Successfully created output File!\nPlease insert correct coordinates into output file and use the push command\n")
		},
	}
	pullCmd.Flags().StringVarP(&output, "output", "o", "coordinates.json", "Flag to specify output file")
	pullCmd.MarkFlagRequired("output")
	return pullCmd
}
