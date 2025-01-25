package migrate

import (
	"fmt"
	"goth/src/config"
	"goth/src/embeded"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
)

type withMigrateFunc func(m *migrate.Migrate)

func apply(fMigrate withMigrateFunc) {
	cfg := config.GetMigrateConfigs()

	d, err := iofs.New(embeded.Migrations, "migrations")
	checkErr("Failed to get embeded migrations", err)

	databaseURL := fmt.Sprintf("sqlite://%s", cfg.DataBaseURL)

	m, err := migrate.NewWithSourceInstance("iofs", d, databaseURL)
	checkErr("Failed to get migrate", err)
	defer m.Close()

	fMigrate(m)
}

func Up(cmd *cobra.Command, args []string) {
	apply(func(m *migrate.Migrate) {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			checkErr("Failed to migrate up", err)
		}
	})
}

func Down(cmd *cobra.Command, args []string) {
	n, err := strconv.Atoi(args[0])
	checkErr("Failed to parse argument", err)

	apply(func(m *migrate.Migrate) {
		if err := m.Steps(-1 * n); err != nil && err != migrate.ErrNoChange {
			checkErr("Failed to migrate down", err)
		}
	})
}

func checkErr(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("[Migrate] %s: %s", msg, err.Error()))
	}
}
