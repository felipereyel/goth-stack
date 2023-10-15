package migrate

import (
	"fmt"
	config "goth/src/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type withMigrateFunc func(m *migrate.Migrate)

func apply(fMigrate withMigrateFunc) {
	sourceURL := fmt.Sprintf("file://%s", config.Configs.MigrationsDir)
	databaseURL := fmt.Sprintf("sqlite://%s", config.Configs.DataBaseURL)

	m, err := migrate.New(sourceURL, databaseURL)
	checkErr("Failed to get migrate", err)
	defer m.Close()

	fMigrate(m)
}

func Up() {
	apply(func(m *migrate.Migrate) {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			checkErr("Failed to migrate up", err)
		}
	})
}

func Down(n int) {
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
