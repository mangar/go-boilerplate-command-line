package cmd

import (
	"fmt"
	"os/exec"
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

	lastLogFileName := getLastFile()
	log.LogDebug("file:"+lastLogFileName, config.FlagDebug)

	//
	cmd := "rm " + lastLogFileName + ".tmp"
	log.LogDebug("cmd: "+cmd, config.FlagDebug)
	exec.Command("sh", "-c", cmd).Run()

	cmd = "tail -n 100 " + lastLogFileName + " > " + lastLogFileName + ".tmp"
	log.LogDebug("cmd: "+cmd, config.FlagDebug)
	exec.Command("sh", "-c", cmd).Run()

	// TODO
	// printa o arquivo...
	cmd = "tail -n 100 " + lastLogFileName + ".tmp"
	log.LogDebug("cmd: "+cmd, config.FlagDebug)
	out, _ := exec.Command("sh", "-c", cmd).Output()

	fmt.Printf("%s", out)

}

func getLastFile() string {
	now := time.Now()
	fileName := config.GetConfig().Log + "/" + now.Format(config.TempFileNameTemplate)
	return fileName
}
