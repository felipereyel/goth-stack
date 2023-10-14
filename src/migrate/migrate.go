package migrate

import (
	"fmt"
	config "goth/src/config"
	"goth/src/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type withMigrateFunc func(m *migrate.Migrate)

func withMigrate(f withMigrateFunc) {
	sourceURL := fmt.Sprintf("file://%s", config.Configs.MigrationsDir)

	db, err := database.New()
	checkErr("Failed to get database", err)
	defer db.Close()

	driver, err := sqlite3.WithInstance(db.Conn, &sqlite3.Config{})
	checkErr("Failed to get driver", err)

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "sqlite3", driver)
	checkErr("Failed to get migrate", err)

	f(m)
}

func Up() {
	withMigrate(func(m *migrate.Migrate) {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			checkErr("Failed to migrate up", err)
		}
	})
}

func Down(n int) {
	withMigrate(func(m *migrate.Migrate) {
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
