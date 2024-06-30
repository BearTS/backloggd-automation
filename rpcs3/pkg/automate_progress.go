package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BearTS/backloggd-automation/pkg/config"
	"github.com/BearTS/backloggd-go/sdk"
	"github.com/bearts/ps3-sfo-parser/sfo"
)

// Searches the saves folder for any saves and adds the details to the Backloggd AI
func AutomateProgress(client *sdk.BackloggdSDK) error {
	// Get savesFolder
	savesFolder, err := config.GetCredentials("rpcs3", "saves-folder")
	if err != nil {
		return err
	}

	saveFiles := []string{}
	// Function to walk through directory and collect .sfo files
	err = filepath.Walk(savesFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".sfo" && strings.ToLower(filepath.Base(path)) == "param.sfo" {
			saveFiles = append(saveFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Read the param.sfo file

	var games []GamesSaveData
	for _, saveFile := range saveFiles {

		var game GamesSaveData
		parser, err := sfo.NewSFOParser(saveFile)
		if err != nil {
			return err
		}

		for i := 0; i < parser.GetLength(); i++ {
			value, err := parser.GetValue("DETAIL")
			if err != nil {
				return err
			}
			// TODO:

			value, err = parser.GetValue("SUB_TITLE")
			if err != nil {
				return err
			}
			game.SubTitle = value.(string)

			value, err = parser.GetValue("TITLE")
			if err != nil {
				return err
			}
			game.Title = value.(string)

			// 			value, err = parser.GetValue("SUB_TITLE")
			// has the string "Progress: 16.9% "
			// get the progress

			progress := strings.Split(game.SubTitle, "Progress: ")
			if len(progress) > 1 {
				floatConversion, err := strconv.ParseFloat(strings.TrimSpace(strings.Split(progress[1], "%")[0]), 64)
				if err != nil {
					return err
				}
				game.Progress = floatConversion
			}

		}
		games = append(games, game)
	}

	fmt.Println(games)

	return nil
}

type GamesSaveData struct {
	Detail   string
	SubTitle string
	Title    string
	Progress float64
}
