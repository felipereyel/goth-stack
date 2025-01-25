package server

import (
	"goth/src/cmd/migrate"
	"goth/src/config"
	"goth/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
)

func Serve(cmd *cobra.Command, args []string) {
	cfg := config.GetServerConfigs()

	if cfg.AutoMigrate {
		migrate.Up(cmd, args)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	app.Use(cors.New())

	if err := routes.Init(app, cfg); err != nil {
		panic(err.Error())
	}

	if err := app.Listen(cfg.ServerAddress); err != nil {
		panic(err.Error())
	}
}
