package routes

import (
	"fmt"
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	fmt.Printf("Route Error [%s]: %v\n", c.Path(), err)
	return sendPage(c, components.ErrorPage())
}

func notFoundHandler(c *fiber.Ctx) error {
	return sendPage(c, components.NotFoundPage())
}
