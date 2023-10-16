package routes

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	app.Use("/healthz", HealthCheck)

	app.Get("/", Home)

	app.Use(NotFound)
}
