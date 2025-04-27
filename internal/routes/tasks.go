package routes

import (
	"goth/internal/components"
	"goth/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func taskList(tc *controllers.TaskController, c *fiber.Ctx) error {
	tasks, err := tc.ListTasks()
	if err != nil {
		return err
	}

	return sendPage(c, components.TaskListPage(tasks))
}

func taskNew(tc *controllers.TaskController, c *fiber.Ctx) error {
	task, err := tc.CreateTask()
	if err != nil {
		return err
	}

	return c.Redirect("/edit/" + task.Id)
}

func taskEdit(tc *controllers.TaskController, c *fiber.Ctx) error {
	taskId := c.Params("id")
	task, err := tc.RetrieveTask(taskId)

	if err != nil {
		return err
	}

	return sendPage(c, components.TaskEditPage(task))
}

func taskSave(tc *controllers.TaskController, c *fiber.Ctx) error {
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
