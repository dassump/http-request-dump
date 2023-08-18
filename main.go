package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/urfave/cli/v2"
)

var version = "devel"

func main() {
	if err := (&cli.App{
		Name:            "http-request-dump",
		Usage:           "HTTP request dump server",
		Description:     "https://github.com/dassump/http-request-dump",
		Authors:         []*cli.Author{{Name: "Daniel Dias de Assumpção", Email: "dassump@gmail.com"}},
		Copyright:       "http://www.apache.org/licenses/LICENSE-2.0",
		Version:         version,
		HideHelpCommand: true,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "listen",
				Usage:   "server address and port",
				Aliases: []string{"l"},
				EnvVars: []string{"HTTPREQUESTDUMP_LISTEN"},
				Value:   "0.0.0.0:8888",
			},
			&cli.BoolFlag{
				Name:    "body",
				Usage:   "dump request body",
				Aliases: []string{"b"},
				EnvVars: []string{"HTTPREQUESTDUMP_BODY"},
				Value:   true,
			},
		},

		Before: func(ctx *cli.Context) error {
			log.SetPrefix(">>> ")
			log.Println("HTTP server listening on", ctx.String("listen"))
			return nil
		},

		Action: func(ctx *cli.Context) error {
			return http.ListenAndServe(ctx.String("listen"), http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					dump, err := httputil.DumpRequest(r, ctx.Bool("body"))
					if err != nil {
						log.Printf("Request from %s\nERROR: %s", r.RemoteAddr, err)
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(err.Error()))
						return
					}

					log.Printf("Request from %s\n%s", r.RemoteAddr, dump)
					w.WriteHeader(http.StatusOK)
					w.Write(dump)
				},
			))
		},
	}).Run(os.Args); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
