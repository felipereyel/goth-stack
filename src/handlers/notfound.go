package handler

import (
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	notFound := components.NotFound()
	return fullPageRender(c, notFound, "Not Found")
}
