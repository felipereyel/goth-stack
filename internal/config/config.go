package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	AutoMigrate   bool
	ServerAddress string

	DataBaseURL string
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	// mandatory

	config.DataBaseURL = os.Getenv("DATABASE_URL")
	if config.DataBaseURL == "" {
		panic("Missing DATABASE_URL")
	}

	// optional - with defaults

	envAutoMigrate := os.Getenv("AUTO_MIGRATE")
	if envAutoMigrate != "" {
		config.AutoMigrate = envAutoMigrate == "true"
	} else {
		config.AutoMigrate = true
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.ServerAddress = ":" + envPort
	} else {
		config.ServerAddress = ":3000"
	}

	return config
}

type MigrateConfigs struct {
	DataBaseURL string
}

func GetMigrateConfigs() MigrateConfigs {
	config := MigrateConfigs{}

	// mandatory

	config.DataBaseURL = os.Getenv("DATABASE_URL")
	if config.DataBaseURL == "" {
		panic("Missing DATABASE_URL")
	}

	return config
}
