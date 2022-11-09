package cmd

import "github.com/spf13/cobra"

var input string // flag to specify input string

func NewPushCmd() *cobra.Command {
	pushCmd := &cobra.Command{
		Use:     "push",
		Short:   "push and save nodes with coordindates from input file into arango ls_node_coordinates collection",
		Example: "gps-injector push -i coordinates.json",

		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	pushCmd.Flags().StringVar(&input, "i", "coordinates.json", "Flag to specify input file")
	pushCmd.MarkFlagRequired("i")
	return pushCmd
}
