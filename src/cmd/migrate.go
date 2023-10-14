package cmd

import (
	"goth/src/migrate"
	"strconv"

	"github.com/spf13/cobra"
)

func migrateUp(cmd *cobra.Command, args []string) {
	migrate.Up()
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Migrates Up the database",
	Run:   migrateUp,
}

func migrateDown(cmd *cobra.Command, args []string) {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		panic("Bad argument: " + err.Error())
	}

	migrate.Down(n)
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Migrates Down the database",
	Run:   migrateDown,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
}
