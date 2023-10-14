package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type tconfigs struct {
	MigrationsDir string
	ServerAddress string
	DataBaseURL   string
	DataBaseName  string
}

func initConfigs() tconfigs {
	var config = tconfigs{}

	envDataBaseURL := os.Getenv("DATABASE_URL")
	if envDataBaseURL != "" {
		config.DataBaseURL = envDataBaseURL
	} else {
		config.DataBaseURL = "db.sqlite3"
	}

	envDataBaseName := os.Getenv("DATABASE_NAME")
	if envDataBaseName != "" {
		config.DataBaseName = envDataBaseName
	} else {
		config.DataBaseName = "goth"
	}

	envMigrationsDir := os.Getenv("MIGRATIONS_DIR")
	if envMigrationsDir != "" {
		config.MigrationsDir = envMigrationsDir
	} else {
		config.MigrationsDir = "migrations"
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.ServerAddress = ":" + envPort
	} else {
		config.ServerAddress = ":3000"
	}

	return config
}

var Configs = initConfigs()
