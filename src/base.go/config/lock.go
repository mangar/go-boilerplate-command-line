package config

import (
	"os"
)

func IsLocked() bool {
	fileName := ConfigDir + "/lock"
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

//
//
//
func Lock() {
	fileName := ConfigDir + "/lock"
	os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
}

//
//
//
func Unlock() {
	fileName := ConfigDir + "/lock"
	os.Remove(fileName)
}
