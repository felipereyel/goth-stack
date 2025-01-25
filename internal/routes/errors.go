package routes

import (
	"fmt"
	"goth/internal/components"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	c.SendStatus(fiber.StatusInternalServerError)
	fmt.Printf("Route Error [%s]: %v\n", c.Path(), err)
	return sendPage(c, components.ErrorPage())
}

func notFoundHandler(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusNotFound)
	return sendPage(c, components.NotFoundPage())
}
