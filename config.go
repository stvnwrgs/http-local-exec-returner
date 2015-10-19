package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config struct
type Config struct {
	BindAddress string
	Certs       Certs
	Server      string
	Paths       []Path
}

// Certs configuration
type Certs struct {
	CaFile   string
	CertFile string
	KeyFile  string
}

// Path with path
type Path struct {
	In              string
	Out             string
	ValidationRegex string
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
