package main

import (
	config "gossr/pkgs/config"
	handler "gossr/pkgs/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views:             engine,
		ViewsLayout:       "layouts/main",
		PassLocalsToViews: true,
		ErrorHandler:      handler.Error,
	})

	app.Use(cors.New())
	initRoutes(app)

	if err := app.Listen(":" + config.Envs.Port); err != nil {
		panic(err.Error())
	}
}

func initRoutes(app *fiber.App) {
	app.Use("/healthz", handler.HealthCheck)
	app.Get("/", handler.Home)
	app.Get("/about", handler.About)
	app.Use(handler.NotFound)
}
