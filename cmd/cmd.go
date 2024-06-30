package cmd

import (
	"fmt"
	"os"

	config "github.com/BearTS/backloggd-automation/config/cmd"
	rpcs3 "github.com/BearTS/backloggd-automation/rpcs3/cmd"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "backlog",
	Short: "backloggd CLI tool to interact with the Backloggd AI",
}

// Execute - starts the CLI
func init() {
	cmd.AddCommand(config.RootCmd)
	cmd.AddCommand(rpcs3.RootCmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
