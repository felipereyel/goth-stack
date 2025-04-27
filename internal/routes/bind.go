package routes

import (
	"goth/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

const cookieName = "goth:jwt"

func controllerBind(controller *controllers.TaskController, handler func(*controllers.TaskController, *fiber.Ctx) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return handler(controller, ctx)
	}
}
