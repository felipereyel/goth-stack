package server

import (
	config "goth/src/config"
	handler "goth/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.Error,
	})

	app.Use(cors.New())
	initRoutes(app)

	if err := app.Listen(config.Configs.ServerAddress); err != nil {
		panic(err.Error())
	}
}

func initRoutes(app *fiber.App) {
	app.Use("/healthz", handler.HealthCheck)
	app.Get("/about", handler.About)
	app.Get("/", handler.Home)
	app.Use(handler.NotFound)
}
