package routes

import (
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func notFoundHandler(c *fiber.Ctx) error {
	c.SendStatus(404)
	notFound := components.NotFound()
	return fullPageRender(c, notFound, "Not Found")
}
