package routes

import (
	"goth/src/components"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	tasks := []components.HomeTask{
		{
			Id:    "dia",
			Title: "Dia",
		},
		{
			Id:    "tarde",
			Title: "Tarde",
		},
		{
			Id:    "noite",
			Title: "Noite",
		},
	}

	home := components.Home(tasks)
	return fullPageRender(c, home, "Home")
}
