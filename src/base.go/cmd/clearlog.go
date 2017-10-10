package cmd

import (
	"os/exec"

	"base.go/config"
	"base.go/log"
)

/**
 *
 * Delete all log files
 *
 */
func ClearLog() {

	logFiles := log.LogFiles()
	for _, fileName := range logFiles {
		log.LogDebug("File:"+config.GetConfig().Log+"/"+fileName, config.FlagDebug)

		if err := exec.Command("sh", "-c", "rm  "+config.GetConfig().Log+"/"+fileName).Run(); err != nil {
			log.LogError("Error removing log file: " + fileName)
		}

	}

}
