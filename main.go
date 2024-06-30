package main

import (
	"github.com/BearTS/backloggd-automation/cmd"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cmd.Execute()
	// 	// REDACTED
}
