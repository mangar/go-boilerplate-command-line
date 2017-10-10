package cmd

import (
	"time"

	"base.go/config"
	"base.go/log"
)

/**
 *
 * Print last thread of execution
 *
 */
func ShowLog() {

	log.LogDebug("file:"+getLastFile(), config.FlagDebug)

}

/**
 * @BOILERPLATE
 *
 * change the log file pattern (go-boilerplate_060201.log) to one that fits your needs
 */
func getLastFile() string {
	now := time.Now()
	fileName := config.GetConfig().Log + "/" + now.Format("go-boilerplate_060201.log")
	return fileName
}
