package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type envs struct {
	PublicDir string
	Port      string
}

var Envs = envs{
	PublicDir: os.Getenv("PUBLIC_DIR"),
	Port:      os.Getenv("PORT"),
}
