package cmd

import (
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

	// abre o arquivo procura o ultimo > > > > >

	// print a partir do ultimo > > > > > até o final do arquivo

}

func getLastFile() string {
	now := time.Now()
	fileName := config.GetConfig().Log + "/" + now.Format(config.TempFileNameTemplate)
	return fileName
}
