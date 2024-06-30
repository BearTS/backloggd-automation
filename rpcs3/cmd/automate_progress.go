package cmd

import (
	"github.com/BearTS/backloggd-automation/rpcs3/pkg"
	"github.com/spf13/cobra"
)

var automateProgress = &cobra.Command{
	Use:   "automate_progress",
	Short: "Automate the progress of a game",
	RunE: func(cmd *cobra.Command, args []string) error {
		// client, err := backloggd.InitClient(false)
		// if err != nil {
		// 	return err
		// }

		err := pkg.AutomateProgress(nil)
		if err != nil {
			return err
		}

		return nil
	},
}
