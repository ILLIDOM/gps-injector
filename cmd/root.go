package cmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gps-injector",
		Short: "CLI",
		Long:  "gps-injector for Jalapeno",
	}

	rootCmd.AddCommand(
		NewPullCmd(),
	)

	return rootCmd
}

func Execute() error {
	return New().Execute()
}
