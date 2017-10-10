package config

import (
	"math/rand"
	"os"
	"os/user"
	"time"

	"base.go/log"
)

var FlagDebug = false

var ConfigDir = HomeDir() + "/.go-boilerplate"
var TempDir = ConfigDir + "/temp"
var ConfigFile = "config.yml"

func Initialize() {
	rand.Seed(time.Now().Unix())

	checkConfigDir()
	// DeleteAndCreateTempDir(constants.ProjectName)
}

func checkConfigDir() {

	_, err := os.Stat(ConfigDir)
	if os.IsNotExist(err) {
		log.LogError("The " + ConfigDir + " dir does not exist")
		os.Exit(1)
	}

	file := ConfigFilePath()
	_, err2 := os.Stat(file)
	if os.IsNotExist(err2) {
		log.LogError("The " + ConfigDir + "/" + ConfigFile + " file does not exist")
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

/**
 * Returns the Config file path (~/.go-boilerplate/config.json)
 */
func ConfigFilePath() string {
	return ConfigDir + "/" + ConfigFile
}

// /**
//  *
//  */
// func DeleteAndCreateTempDir(project string) {

// 	tempDir := HomeDir() + "/.gonfig/temp/" + project
// 	log.LogDebug("Cleanning temporary dir @" + tempDir)

// 	if err := exec.Command("sh", "-c", "rm -rf "+tempDir).Run(); err != nil {
// 		log.LogError("Error removing build directory")
// 	}

// 	CreateTempDir(project)
// }

// /**
//  *
//  */
// func CreateTempDir(project string) {

// 	tempDir := HomeDir() + "/.gonfig/temp/" + project

// 	_, err := os.Stat(tempDir)
// 	if os.IsNotExist(err) {
// 		log.LogDebug("Creating temp dir @" + tempDir)
// 		os.MkdirAll(tempDir, os.ModePerm)
// 	}
// }
