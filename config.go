package main

import (
	"encoding/json"
	"os"
)

func readConf() error {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}
	if json.Unmarshal(configFile, &config) != nil {
		return err
	}
	return nil
}
