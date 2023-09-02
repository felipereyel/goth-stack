package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	return c.Render("error", fiber.Map{
		"Error": err.Error(),
		"Title": "Internal Error",
	})
}
