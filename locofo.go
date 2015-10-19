package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func buildTlsClient(certs Certs) *http.Client {
	// Load client cert
	cert, err := tls.LoadX509KeyPair(certs.CertFile, certs.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(certs.CaFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	return client
}

func request(config Config, path Config.Paths) {
	// Load client cert
	client := buildTlsClient(config.Certs)

	// Do GET something
	resp, err := client.Get("https://goldportugal.local:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}

func requestHandler(config Config, out Path.Out) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		out, errCode := request(config, path)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		io.WriteString(w, string(out[:]))
	}
}

// Serve runs the Webserver
func serveHttp(config Config) {
	mux := http.NewServeMux()

	for i := range config.Paths {
		path := config.Paths[i]
		mux.HandleFunc(path.In, requestHandler(config, path.Out))
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
