package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "config",
	Short: "Change the configuration of the CLI",
}

// Execute - starts the CLI
func init() {
	RootCmd.AddCommand(rpcs3Config)
	RootCmd.AddCommand(backloggdConfig)
}
