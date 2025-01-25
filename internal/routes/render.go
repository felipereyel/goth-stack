package routes

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func sendPage(c *fiber.Ctx, page templ.Component) error {
	c.Set("Content-Type", "text/html")
	return page.Render(context.Background(), c)
}
