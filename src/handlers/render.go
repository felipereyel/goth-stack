package handler

import (
	"context"
	"goth/src/components"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func fullPageRender(c *fiber.Ctx, body templ.Component, title string) error {
	c.Set("Content-Type", "text/html")
	return components.FullPage(title, body).Render(context.Background(), c)
}
