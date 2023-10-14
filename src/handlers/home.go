package handler

import (
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	home := components.Home("Welcome to my website")
	return fullPageRender(c, home, "Home")
}
