package serve

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var port string

func NewserveCmd() (serveCmd *cobra.Command) {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Create a HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.WithFields(log.Fields{
					"url":     r.URL,
					"method":  r.Method,
					"headers": r.Header,
				}).Infoln("Incoming HTTP Request")

				w.WriteHeader(http.StatusNoContent)
			})

			log.Fatal(http.ListenAndServe(":"+port, handlerFunc))
		},
	}

	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "80", "HTTP server exposed port")

	return serveCmd
}
