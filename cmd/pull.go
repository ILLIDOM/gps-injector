package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var output string // flag to specify output file

func NewPullCmd() *cobra.Command {
	pullCmd := &cobra.Command{
		Use:     "pull",
		Short:   "get and save yaml for coordinates",
		Example: "gps-injector pull -o coordinates.yaml",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pull command executed")
		},
	}
	pullCmd.Flags().StringVar(&output, "-o", "coordinates.yaml", "Flag to specify output file")
	return pullCmd
}
