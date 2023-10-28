package cmd

import (
	"goth/src/cmd/fiddle"
	"goth/src/cmd/migrate"
	"goth/src/cmd/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goth",
	Short: "goth app CLI",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Run:   server.Start,
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Migrates Up the database",
	Run:   migrate.Up,
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Migrates Down the database",
	Run:   migrate.Down,
	Args:  cobra.ExactArgs(1),
}

var fiddleCmd = &cobra.Command{
	Use:   "fiddle",
	Short: "Fiddle with the application",
	Run:   fiddle.Run,
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(fiddleCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
