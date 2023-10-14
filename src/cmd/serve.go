package cmd

import (
	"goth/src/server"

	"github.com/spf13/cobra"
)

func serve(cmd *cobra.Command, args []string) {
	server.Start()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
