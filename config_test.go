package main

import (
	"testing"
)

var expectedBindAdress = "0.0.0.0:2350"
var expectedCommand = "echo"

func TestLoadConfig(t *testing.T) {
	conf := LoadConfig("./exampleconf.json")
	if conf.BindAddress != expectedBindAdress {
		t.Errorf("Expected: %s | But got: %s", expectedBindAdress, conf.BindAddress)
	}
	if conf.Commands[0].Command != expectedCommand {
		t.Errorf("Expected: %s | But got: %s", expectedCommand, conf.Commands[0].Command)
	}
}
