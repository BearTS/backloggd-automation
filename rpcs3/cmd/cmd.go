package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "rpcs3",
	Short: "RPCS3 tool to interact with the RPCS3 emulator and backloggd API",
}

// Execute - starts the CLI
func init() {
	RootCmd.AddCommand(automateProgress)
}
