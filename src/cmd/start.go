package cmd

import (
	"goth/src/migrate"
	"goth/src/server"

	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, args []string) {
	migrate.Up()
	server.Start()
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Migration Up and Start the server",
	Run:   start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}
