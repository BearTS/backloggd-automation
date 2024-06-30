package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BearTS/backloggd-automation/pkg/utils"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var configName = "config"
var fileName = configName + ".toml"
var (
	k      = koanf.New(".")
	parser = toml.Parser()
)

func init() {
	// Create the directory and config file if they don't exist
	if err := ensureConfigFile(); err != nil {
		panic(fmt.Errorf("failed to initialize config file: %s", err))
	}
	// Set the path to the directory containing the config file
	k.Load(file.Provider(filepath.Join(utils.GetConfigDirectory())+"/"+fileName), parser)
}

func ensureConfigFile() error {
	// Get the path to the config file
	configPath := filepath.Join(utils.GetConfigDirectory(), fileName)

	// Check if the config file exists
	if !filePathExists(configPath) {
		// Config file does not exist, create the directory and file
		err := os.MkdirAll(filepath.Dir(configPath), 0700)
		if err != nil {
			return err
		}

		_, err = os.Create(configPath)
		if err != nil {
			return err
		}

		fmt.Println("Created config file:", configPath)
	}
	return nil
}

func SetCredentials(appName string, key string, value string) {
	config := k.Get(appName)
	if config == nil {
		config = make(map[string]interface{})
	}

	appConfig := config.(map[string]interface{})

	if _, ok := appConfig[key]; ok {
		// Key already exists, replace the value
		appConfig[key] = value
	} else {
		// Key doesn't exist, add a new key-value pair
		appConfig[key] = value
	}

	err := k.Set(appName, appConfig)
	if err != nil {
		panic(fmt.Errorf("failed to set credentials: %s", err))
	}

	b, err := k.Marshal(parser)
	if err != nil {
		panic(fmt.Errorf("failed to marshal config: %s", err))
	}

	// Write back to the config file
	configPath := filepath.Join(utils.GetConfigDirectory(), fileName)
	err = os.WriteFile(configPath, b, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to write to config file: %s", err))
	}
}

func GetCredentials(appName string, key string) (string, error) {
	appConfig := k.Get(appName)
	if appConfig == nil {
		return "", fmt.Errorf("app %s not found", appName)
	}

	appConfigMap := appConfig.(map[string]interface{})
	if value, ok := appConfigMap[key]; ok {
		return value.(string), nil
	}

	return "", fmt.Errorf("key %s not found in app %s", key, appName)
}

func filePathExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist
			return false
		} // Some other error occurred
		panic(fmt.Errorf("failed to check if file exists: %s", err))
	}
	return true
}
