package utils

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BearTS/backloggd-automation/pkg/constants"
)

func GetConfigDirectory() string {
	return filepath.Join(homeDir(), constants.ConfigDirectory)
}

// homeDir returns the path to the user's home directory.
func homeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}
