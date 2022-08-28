package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
)

var (
	listen         string
	listen_key     = "listen"
	listen_default = "0.0.0.0:8888"
	listen_info    = "Server address and port"
)

func init() {
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
