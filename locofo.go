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
	"regexp"
)

func isValid(text string, regex string) bool {
	match, _ := regexp.MatchString(regex, text)
	log.Printf("Response body matching %s with %s ", regex, text)
	log.Printf("Response body regex match result: %t ", match)
	return match
}

func buildTLSClient(certs Certs) *http.Client {
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

func request(client *http.Client, uri string) (string, int) {
	// Do GET something
	statusCode := 200

	resp, err := client.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		statusCode = resp.StatusCode
	}

	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data), statusCode
}

func requestHandler(config Config, path Path) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		uri := config.Server + path.Out
		tlsClient := buildTLSClient(config.Certs)
		responseBody, statusCode := request(tlsClient, uri)

		if statusCode == 200 {
			log.Printf("Successfull [200] response with body : %s", responseBody)
			if !isValid(responseBody, path.ValidationRegex) {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		io.WriteString(w, string(responseBody[:]))
	}
}

// Serve runs the Webserver
func serveHttp(config Config) {
	mux := http.NewServeMux()

	for i := range config.Paths {
		path := config.Paths[i]
		mux.HandleFunc(path.In, requestHandler(config, path))
	}

	log.Fatalf("Error starting the http service:%s", http.ListenAndServe(config.BindAddress, mux))
}

func main() {
	args := os.Args
	if (len(args) <= 1) && ((args[1] == "-h") || (args[1] == "--help")) {
		fmt.Println("Usage: \n locofo <config_file>")
		os.Exit(2)
	}
	config = LoadConfig(args[1])

	serveHttp(config)
}
