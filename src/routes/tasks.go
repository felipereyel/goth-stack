package routes

import (
	"goth/src/components"
	"goth/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func initTaskRoutes(app *fiber.App, tc *controllers.TaskController) {
	app.Get("/", taskList(tc))
	app.Get("/new", taskNew(tc))
	app.Get("/edit/:id", taskEdit(tc))
	app.Post("/edit/:id", taskSave(tc))
}

func taskList(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mockOwner := "00000000-0000-4000-0000-000000000000"
		tasks, err := tc.ListTasks(mockOwner)

		if err != nil {
			return err
		}

		return sendPage(c, components.TaskListPage(tasks))
	}
}

func taskNew(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mockOwner := "00000000-0000-4000-0000-000000000000"
		task, err := tc.CreateTask(mockOwner)

		if err != nil {
			return err
		}

		return c.Redirect("/edit/" + task.Id)
	}
}

func taskEdit(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		taskId := c.Params("id")
		task, err := tc.RetrieveTask(taskId)

		if err != nil {
			return err
		}

		return sendPage(c, components.TaskEditPage(task))
	}
}

func taskSave(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var taskId = c.Params("id")
		var taskChange controllers.TaskChange
		err := c.BodyParser(&taskChange)
		if err != nil {
			return err
		}

		if err := tc.UpdateTask(taskId, taskChange); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
