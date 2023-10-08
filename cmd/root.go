package cmd

import (
	"os"

	"github.com/emmanuelgautier/httpmock-netlogger/cmd/serve"
	"github.com/spf13/cobra"
)

func NewRootCmd() (cmd *cobra.Command) {
	var rootCmd = &cobra.Command{
		Use:   "httpmock-netlogger",
		Short: "A simple HTTP Server mock which log every network connections events.",
	}
	rootCmd.AddCommand(serve.NewserveCmd())

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
