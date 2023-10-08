package serve

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewserveCmd() (serveCmd *cobra.Command) {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Create an HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.WithFields(log.Fields{
					"url":     r.URL,
					"method":  r.Method,
					"headers": r.Header,
				}).Info("Incoming HTTP Request")

				w.WriteHeader(http.StatusNoContent)
			})

			log.Fatal(http.ListenAndServe(":8080", handlerFunc))
		},
	}

	return serveCmd
}
