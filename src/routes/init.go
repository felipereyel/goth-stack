package routes

import (
	"fmt"
	"goth/src/config"
	"goth/src/controllers"
	"goth/src/repositories/database"
	"goth/src/repositories/jwt"
	"goth/src/repositories/oidc"

	"github.com/gofiber/fiber/v2"
)

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Init(app *fiber.App, cfg config.ServerConfigs) error {
	jwtRepo := jwt.NewJWTRepo(cfg.JwtSecret)

	dbRepo, err := database.NewDatabaseRepo(cfg.DataBaseURL)
	if err != nil {
		return fmt.Errorf("[Init] failed to get database: %w", err)
	}

	oidcRepo, err := oidc.NewOIDC(
		cfg.OIDCIssuer,
		cfg.OIDCClientID,
		cfg.OIDCClientSec,
		cfg.OIDCRedirectURI,
	)
	if err != nil {
		return fmt.Errorf("[Init] failed to get oidc: %w", err)
	}

	uc := controllers.NewUserController(dbRepo, oidcRepo, jwtRepo)
	tc := controllers.NewTaskController(dbRepo, oidcRepo)

	app.Get("/auth/redirect", authRedirect(uc))
	app.Get("/auth/login", loginHandler(uc))

	app.Get("/", verifyAuth(uc), taskList(tc))
	app.Get("/new", verifyAuth(uc), taskNew(tc))
	app.Get("/edit/:id", verifyAuth(uc), taskEdit(tc))
	app.Post("/edit/:id", verifyAuth(uc), taskSave(tc))

	app.Use("/healthz", healthCheckHandler)
	app.Use(notFoundHandler)

	return nil
}
