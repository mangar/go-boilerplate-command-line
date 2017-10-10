package cmd

import (
	"fmt"

	"base.go/config"
)

/**
 * @BOILERPLATE
 *
 * Replace the content from config file you want to be presented
 * Sometimes just the file path is enought.
 *
 */
func ShowConfig() {
	config.GetConfig()

	fmt.Println("Configuration: (", config.ConfigDir+"/"+config.ConfigFile, ")")
	fmt.Println("")
	fmt.Println("Log:", config.Config.Log)

}
