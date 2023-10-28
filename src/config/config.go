package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	AutoMigrate   bool
	MigrationsDir string
	ServerAddress string

	DataBaseURL string
	JwtSecret   string

	OIDCIssuer            string
	OIDCClientID          string
	OIDCClientSec         string
	OIDCRedirectURI       string
	OIDCLogoutRedirectURI string
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	// mandatory

	config.DataBaseURL = os.Getenv("DATABASE_URL")
	if config.DataBaseURL == "" {
		panic("Missing DATABASE_URL")
	}

	config.JwtSecret = os.Getenv("JWT_SECRET")
	if config.JwtSecret == "" {
		panic("Missing JWT_SECRET")
	}

	config.OIDCIssuer = os.Getenv("OIDC_ISSUER")
	if config.OIDCIssuer == "" {
		panic("Missing OIDC_ISSUER")
	}

	config.OIDCClientID = os.Getenv("OIDC_CLIENT_ID")
	if config.OIDCClientID == "" {
		panic("Missing OIDC_CLIENT_ID")
	}

	config.OIDCClientSec = os.Getenv("OIDC_CLIENT_SECRET")
	if config.OIDCClientSec == "" {
		panic("Missing OIDC_CLIENT_SECRET")
	}

	config.OIDCRedirectURI = os.Getenv("OIDC_REDIRECT_URI")
	if config.OIDCRedirectURI == "" {
		panic("Missing OIDC_REDIRECT_URI")
	}

	config.OIDCLogoutRedirectURI = os.Getenv("OIDC_LOGOUT_REDIRECT_URI")
	if config.OIDCLogoutRedirectURI == "" {
		panic("Missing OIDC_LOGOUT_REDIRECT_URI")
	}

	// optional - with defaults

	envAutoMigrate := os.Getenv("AUTO_MIGRATE")
	if envAutoMigrate != "" {
		config.AutoMigrate = envAutoMigrate == "true"
	} else {
		config.AutoMigrate = true
	}

	config.MigrationsDir = os.Getenv("MIGRATIONS_DIR")
	if config.MigrationsDir == "" {
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

type MigrateConfigs struct {
	MigrationsDir string
	DataBaseURL   string
}

func GetMigrateConfigs() MigrateConfigs {
	config := MigrateConfigs{}

	// mandatory

	config.DataBaseURL = os.Getenv("DATABASE_URL")
	if config.DataBaseURL == "" {
		panic("Missing DATABASE_URL")
	}

	// optional - with defaults

	config.MigrationsDir = os.Getenv("MIGRATIONS_DIR")
	if config.MigrationsDir == "" {
		config.MigrationsDir = "migrations"
	}

	return config
}
