package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gps-injector",
		Short: "CLI",
		Long:  "gps-injector for Jalapeno",
	}

	rootCmd.AddCommand(
		NewPullCmd(),
		NewPushCmd(),
	)

	return rootCmd
}

func Execute() error {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	return New().Execute()
}
