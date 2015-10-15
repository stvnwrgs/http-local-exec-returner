package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config struct
type Config struct {
	BindAddress string
	Commands    []Command
}

// Command with path
type Command struct {
	Path    string
	Command string
}

var config = Config{}

func LoadConfig(path string) Config {
	configFile, err := os.Open(path)
	if err != nil {
		fmt.Errorf("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Errorf("parsing config file", err.Error())
	}

	return config
}
