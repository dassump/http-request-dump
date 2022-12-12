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
	app         = "http-request-dump"
	version     = "unknown"
	description = "HTTP request dump server"
	site        = "https://github.com/dassump/http-request-dump"

	listen = flag.String("listen", "0.0.0.0:8888", "Server address and port")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s (%s)\n\n%s\n%s\n\n", app, version, description, site)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	log.SetPrefix(">>> ")
	log.Println("Listening on", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			dump, _ := httputil.DumpRequest(r, true)
			log.Printf("Request from %s\n%s", r.RemoteAddr, dump)
			w.WriteHeader(http.StatusOK)
		},
	)))
}
