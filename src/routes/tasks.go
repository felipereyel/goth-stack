package routes

import (
	"goth/src/components"
	"goth/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func taskList(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			return err
		}

		tasks, err := tc.ListTasks(user.ID)
		if err != nil {
			return err
		}

		return sendPage(c, components.TaskListPage(tasks))
	}
}

func taskNew(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			return err
		}

		task, err := tc.CreateTask(user.ID)
		if err != nil {
			return err
		}

		return c.Redirect("/edit/" + task.Id)
	}
}

func taskEdit(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			return err
		}

		taskId := c.Params("id")
		task, err := tc.RetrieveTask(user.ID, taskId)

		if err != nil {
			return err
		}

		return sendPage(c, components.TaskEditPage(task))
	}
}

func taskSave(tc *controllers.TaskController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := getUser(c)
		if err != nil {
			return err
		}

		var taskId = c.Params("id")
		var taskChange controllers.TaskChange
		err = c.BodyParser(&taskChange)
		if err != nil {
			return err
		}

		if err := tc.UpdateTask(user.ID, taskId, taskChange); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
