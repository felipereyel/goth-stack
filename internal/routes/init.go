package routes

import (
	"fmt"
	"goth/internal/config"
	"goth/internal/controllers"
	"goth/internal/repositories/database"

	"github.com/gofiber/fiber/v2"
)

func healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Init(app *fiber.App, cfg config.ServerConfigs) error {
	dbRepo, err := database.NewDatabaseRepo(cfg)
	if err != nil {
		return fmt.Errorf("[Init] failed to get database: %w", err)
	}

	tc := controllers.NewTaskController(dbRepo)

	app.Get("/", controllerBind(tc, taskList))
	app.Get("/new", controllerBind(tc, taskNew))
	app.Get("/edit/:id", controllerBind(tc, taskEdit))
	app.Post("/edit/:id", controllerBind(tc, taskSave))

	app.Use("/statics", staticsHandler)
	app.Use("/healthz", healthzHandler)
	app.Use(notFoundHandler)

	return nil
}
