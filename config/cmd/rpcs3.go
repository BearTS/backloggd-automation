package cmd

import (
	"fmt"
	"os"

	"github.com/BearTS/backloggd-automation/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rpcs3Config = &cobra.Command{
	Use:   "rpcs3",
	Short: "Change the RPCS3 games folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve flag values
		gamesFolder, err := cmd.Flags().GetString("games-folder")
		if err != nil {
			return err
		}
		savesFolder, err := cmd.Flags().GetString("saves-folder")
		if err != nil {
			return err
		}

		// Handle the logic for games-folder and saves-folder
		if gamesFolder != "" {
			// Try to read the folder
			fmt.Printf("Setting games folder to: %s\n", gamesFolder)
			files, err := os.ReadDir(gamesFolder)
			if err != nil {
				return fmt.Errorf("failed to read games folder: %s", err)
			}

			if len(files) == 0 {
				log.Warn().Msg("No games found in the folder")
			}
			config.SetCredentials("rpcs3", "games-folder", gamesFolder)

		}

		if savesFolder != "" {
			fmt.Printf("Setting saves folder to: %s\n", savesFolder)
			files, err := os.ReadDir(savesFolder)
			if err != nil {
				return fmt.Errorf("failed to read saves folder: %s", err)
			}

			if len(files) == 0 {
				log.Warn().Msg("No saves found in the folder")
			}

			config.SetCredentials("rpcs3", "saves-folder", savesFolder)
		}

		return nil
	},
}

func init() {
	// Add persistent flags for games-folder and saves-folder
	rpcs3Config.PersistentFlags().String("games-folder", "", "Absolute Folder Path for RPCS3 games")
	rpcs3Config.PersistentFlags().String("saves-folder", "", "Absolute Folder Path for RPCS3 game saves")
}