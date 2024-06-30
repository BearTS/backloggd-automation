package cmd

import (
	"fmt"

	"github.com/BearTS/backloggd-automation/pkg/backloggd"
	"github.com/BearTS/backloggd-automation/pkg/config"
	"github.com/spf13/cobra"
)

var backloggdConfig = &cobra.Command{
	Use:   "backloggd",
	Short: "Change the configuration of the Backloggd AI",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve flag values
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		// Handle the logic for username and password
		if username != "" {
			// Try to read the folder
			fmt.Printf("Setting username to: %s\n", username)
			config.SetCredentials("backloggd", "username", username)
		}

		if password != "" {
			fmt.Printf("Changing password\n")
			config.SetCredentials("backloggd", "password", password)
		}

		_, err = backloggd.InitClient(true)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	backloggdConfig.Flags().String("username", "", "The username for the Backloggd AI")
	backloggdConfig.Flags().String("password", "", "The password for the Backloggd AI")
}
