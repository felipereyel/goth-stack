package routes

import (
	"fmt"
	"goth/internal/config"
	"goth/internal/controllers"
	"goth/internal/repositories/database"
	"goth/internal/repositories/jwt"
	"goth/internal/repositories/oidc"

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

	oidcRepo, err := oidc.NewOIDC(cfg)
	if err != nil {
		return fmt.Errorf("[Init] failed to get oidc: %w", err)
	}

	uc := controllers.NewUserController(dbRepo, oidcRepo, jwtRepo)
	tc := controllers.NewTaskController(dbRepo, oidcRepo)

	app.Get("/auth/login", loginHandler(uc))
	app.Get("/auth/post-login", postLoginHandler(uc))

	app.Get("/auth/logout", logoutHandler(uc))
	app.Get("/auth/post-logout", postLogoutHandler(uc))

	app.Get("/", withAuth(uc, tc, taskList))
	app.Get("/new", withAuth(uc, tc, taskNew))
	app.Get("/edit/:id", withAuth(uc, tc, taskEdit))
	app.Post("/edit/:id", withAuth(uc, tc, taskSave))

	app.Use("/statics", staticsHandler)
	app.Use("/healthz", healthzHandler)
	app.Use(notFoundHandler)

	return nil
}
