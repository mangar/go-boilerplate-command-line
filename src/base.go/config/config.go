package config

import (
	"encoding/json"
	"io/ioutil"
)

/**
 * @BOILERPLATE
 *
 * structure of the config file
 */
type ConfigSetup struct {
	Log string `json:log`
}

var Config = ConfigSetup{}

func GetConfig() ConfigSetup {
	data, _ := ioutil.ReadFile(ConfigDir + "/" + ConfigFile)
	json.Unmarshal([]byte(data), &Config)

	return Config
}
