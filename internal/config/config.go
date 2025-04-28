package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	IsProd bool
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	config.IsProd = true
	if os.Getenv("ENV") == "local" {
		config.IsProd = false
	}

	return config
}
