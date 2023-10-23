package routes

import (
	"fmt"
	"goth/src/controllers"
	"goth/src/repositories/database"

	"github.com/gofiber/fiber/v2"
)

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Init(app *fiber.App) error {
	database, err := database.NewDatabaseRepo()
	if err != nil {
		return fmt.Errorf("[Init] failed to get database: %w", err)
	}

	tc := controllers.NewTaskController(database)
	initTaskRoutes(app, tc)

	app.Use("/healthz", healthCheckHandler)
	app.Use(notFoundHandler)

	return nil
}
