package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Excute runs a command and returns output []byte and err
func Execute(command string, args string) (stringValue []byte, err error) {
	cmd := exec.Command(command, args)
	output, err := cmd.CombinedOutput()
	fmt.Printf("Command executed: %s\n", strings.Join(cmd.Args, " "))

	if output != nil {
		fmt.Printf("Command Stdout: %s\n", string(output[:]))
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	return output, err
}
