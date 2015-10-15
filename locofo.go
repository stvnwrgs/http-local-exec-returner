package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(command Command) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		out, err := Execute(command.Command, command.Args)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		io.WriteString(w, string(out[:]))
	}
}

// Serve runs the Webserver
func Serve(config Config) {
	mux := http.NewServeMux()

	for i := range config.Commands {
		c := config.Commands[i]
		mux.HandleFunc(c.Path, handler(c))
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
