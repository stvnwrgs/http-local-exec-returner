package main

import (
	"encoding/json"
	"log"
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
		log.Fatalf("Error opening config file %s", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatalf("Error parsing config file %s", err.Error())
	}

	return config
}
