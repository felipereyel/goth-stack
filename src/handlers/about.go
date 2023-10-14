package handler

import (
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func About(c *fiber.Ctx) error {
	about := components.About("About My Website")
	return fullPageRender(c, about, "About")
}
