package setup

import (
	"math/rand"
	"os"
	"os/user"
	"time"

	"base.go/config"
	"base.go/log"
)

func Initialize() {
	rand.Seed(time.Now().Unix())
	checkConfigDir()
}

func checkConfigDir() {

	_, err := os.Stat(config.ConfigDir)
	if os.IsNotExist(err) {
		log.LogError("The " + config.ConfigDir + " dir does not exist")
		os.Exit(1)
	}

	_, err2 := os.Stat(config.ConfigDir + "/" + config.ConfigFile)
	if os.IsNotExist(err2) {
		log.LogError("The " + config.ConfigDir + "/" + config.ConfigFile + " file does not exist")
		os.Exit(1)
	}

}

/**
 * Returns the HomeDir of the logged in user
 */
func HomeDir() string {
	var homeDir string
	usr, err := user.Current()

	if err == nil {
		homeDir = usr.HomeDir

	} else {
		homeDir = os.Getenv("HOME")
	}

	return homeDir
}
