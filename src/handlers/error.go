package handler

import (
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	errorComponent := components.Error(err.Error())
	return fullPageRender(c, errorComponent, "Internal Error")
}
