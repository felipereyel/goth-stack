package routes

import (
	"fmt"
	"goth/internal/config"
	"goth/internal/controllers"
	"goth/internal/repositories/database"
	"goth/internal/repositories/jwt"

	"github.com/gofiber/fiber/v2"
)

func healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Init(app *fiber.App, cfg config.ServerConfigs) error {
	jwtRepo := jwt.NewJWTRepo(cfg)

	dbRepo, err := database.NewDatabaseRepo(cfg)
	if err != nil {
		return fmt.Errorf("[Init] failed to get database: %w", err)
	}

	uc := controllers.NewUserController(dbRepo, jwtRepo)
	tc := controllers.NewTaskController(dbRepo)

	app.Get("/auth/login", getLoginHandler)
	app.Post("/auth/login", postLoginHandler(uc))

	app.Get("/auth/register", getRegisterHandler)
	app.Post("/auth/register", postRegisterHandler(uc))

	app.Get("/auth/logout", getLogoutHandler)

	app.Get("/", withAuth(uc, tc, taskList))
	app.Get("/new", withAuth(uc, tc, taskNew))
	app.Get("/edit/:id", withAuth(uc, tc, taskEdit))
	app.Post("/edit/:id", withAuth(uc, tc, taskSave))

	app.Use("/statics", staticsHandler)
	app.Use("/healthz", healthzHandler)
	app.Use(notFoundHandler)

	return nil
}
