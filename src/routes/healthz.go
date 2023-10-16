package routes

import (
	"github.com/gofiber/fiber/v2"
)

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
