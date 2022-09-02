package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	app         string = "http-request-dump"
	version     string = "dev"
	description string = "HTTP request dump server"
	site        string = "https://github.com/dassump/http-request-dump"
	info        string = "%s (%s)\n\n%s\n%s\n\n"
	usage       string = "Usage of %s:\n"

	listen         string
	listen_key     = "listen"
	listen_default = "0.0.0.0:8888"
	listen_info    = "Server address and port"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), info, app, version, description, site)
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&listen, listen_key, listen_default, listen_info)
	flag.Parse()
}

func main() {
	log.SetPrefix(">>> ")
	log.Println("Listening on", listen)
	log.Fatal(http.ListenAndServe(listen, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			dump, _ := httputil.DumpRequest(r, true)
			log.Printf("Request from %s\n%s", r.RemoteAddr, dump)
			w.WriteHeader(http.StatusOK)
		},
	)))
}
