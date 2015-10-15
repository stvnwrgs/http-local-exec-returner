package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func exec(command string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!"+command)
	}
}

func Serve(config Config) {
	mux := http.NewServeMux()

	for i := range config.Commands {
		c := config.Commands[i]
		mux.HandleFunc(c.Path, exec(c.Command))
	}

	http.ListenAndServe(config.BindAddress, mux)
}

func main() {

	args := os.Args[1:]
	if (len(args) != 1) || (args[0] == "-h") || (args[0] == "--help") {
		fmt.Print("Usage: \n locofo <config_file>")
	}
	config = LoadConfig(args[0])

	Serve(config)
}
