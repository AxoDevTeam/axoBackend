package config

import (
	"encoding/json"
	"os"
)

var Conf map[string]interface{}

func ReadConf() error {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}
	if json.Unmarshal(configFile, &Conf) != nil {
		return err
	}
	return nil
}
