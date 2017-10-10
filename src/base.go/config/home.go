package config

import (
	"os"
	"os/user"
	"strconv"
)

var ProcessID = strconv.Itoa(os.Getpid())

var FlagDebug = false
var WriteLog = true

var ConfigDir = homeDir() + "/.go-boilerplate"
var TempDir = ConfigDir + "/temp"
var ConfigFile = "config.yml"

func homeDir() string {
	var homeDir string
	usr, err := user.Current()

	if err == nil {
		homeDir = usr.HomeDir

	} else {
		homeDir = os.Getenv("HOME")
	}

	return homeDir
}
