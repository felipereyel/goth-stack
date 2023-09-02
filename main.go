package main

import (
	config "gossr/pkgs/config"
	handler "gossr/pkgs/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use("/healthz", handler.HealthCheck)

	err := app.Listen(":" + config.Envs.Port)
	if err != nil {
		panic(err.Error())
	}
}
