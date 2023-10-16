package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type tconfigs struct {
	AutoMigrate   bool
	MigrationsDir string
	ServerAddress string
	DataBaseURL   string
	DataBaseName  string
}

var Configs tconfigs

func init() {
	envDataBaseURL := os.Getenv("DATABASE_URL")
	if envDataBaseURL != "" {
		Configs.DataBaseURL = envDataBaseURL
	} else {
		Configs.DataBaseURL = "db.sqlite"
	}

	envDataBaseName := os.Getenv("DATABASE_NAME")
	if envDataBaseName != "" {
		Configs.DataBaseName = envDataBaseName
	} else {
		Configs.DataBaseName = "goth"
	}

	envAutoMigrate := os.Getenv("AUTO_MIGRATE")
	if envAutoMigrate != "" {
		Configs.AutoMigrate = envAutoMigrate == "true"
	} else {
		Configs.AutoMigrate = true
	}

	envMigrationsDir := os.Getenv("MIGRATIONS_DIR")
	if envMigrationsDir != "" {
		Configs.MigrationsDir = envMigrationsDir
	} else {
		Configs.MigrationsDir = "migrations"
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		Configs.ServerAddress = ":" + envPort
	} else {
		Configs.ServerAddress = ":3000"
	}
}
