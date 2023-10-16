package routes

import (
	"fmt"
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	fmt.Printf("Route Error [%s]: %v\n", c.Path(), err)
	errorComponent := components.Error()
	return fullPageRender(c, errorComponent, "Internal Error")
}
