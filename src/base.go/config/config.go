package config

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigSetup struct {
	Log string `json:log`
}

var Config = ConfigSetup{}

func GetConfig() ConfigSetup {
	data, _ := ioutil.ReadFile(ConfigFilePath())
	json.Unmarshal([]byte(data), &Config)

	// fmt.Println(&Config)
	return Config
}
