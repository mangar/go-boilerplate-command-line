package cmd

import (
	"io/ioutil"
	"os"

	"base.go/config"
	"base.go/log"
)

/**
 *
 * The installation process is matter of:
 *  1) create the project dir inside the ~
 *  2) create the basic config.json file and
 *  3) create the temporary dir inside the project one
 *
 */
func Install() {

	createHome()
	createConfigFile()
	createTempDir()
	log.LogSuccess("Installation complete!")

}

func createHome() {

	tempDir := config.ConfigDir

	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		log.LogDebug("Creating home dir @"+tempDir, config.FlagDebug)
		os.MkdirAll(tempDir, os.ModePerm)
	} else {
		log.LogDebug("Home dir already created. Nothing to do", config.FlagDebug)
	}

}

func createTempDir() {

	tempDir := config.TempDir

	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		log.LogDebug("Creating temp dir @"+tempDir, config.FlagDebug)
		os.MkdirAll(tempDir, os.ModePerm)
	} else {
		log.LogDebug("Temp dir already created. Nothing to do", config.FlagDebug)
	}

}

/*
 * @BOILERPLATE
 *
 * The base content for config.json file
 *
 */
func createConfigFile() {

	tempDir := config.ConfigFilePath()
	_, err := os.Stat(tempDir)

	if os.IsNotExist(err) {
		log.LogDebug("Creating config.json @"+tempDir, config.FlagDebug)

		_content := []byte(
			"{\n" +
				"	\"log\":\"/var/tmp\"\n" +
				"}\n")

		err := ioutil.WriteFile(tempDir, _content, 0644)
		if err != nil {
			log.LogError("Error writing on file")
		}

	} else {
		log.LogDebug("Config file already there. Nothing to do", config.FlagDebug)
	}

}
