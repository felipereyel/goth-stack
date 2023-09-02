package handler

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	return c.Render("notfound", fiber.Map{
		"Title": "Not Found",
	})
}
