package serve

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func NewserveCmd() (serveCmd *cobra.Command) {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Create an HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.Println(r)

				w.WriteHeader(http.StatusNoContent)
			})

			log.Fatal(http.ListenAndServe(":8080", handlerFunc))
		},
	}

	return serveCmd
}
