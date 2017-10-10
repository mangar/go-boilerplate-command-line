package cmd

import (
	"fmt"

	"base.go/config"
)

func ShowConfig() {
	config.GetConfig()

	fmt.Println("Configuration: (", config.ConfigFilePath(), ")")
	fmt.Println("")
	fmt.Println("Log:", config.Config.Log)

}
