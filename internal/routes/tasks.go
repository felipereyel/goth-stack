package routes

import (
	"errors"
	"goth/internal/components"
	"goth/internal/models"
	"goth/internal/repositories/database"

	"github.com/pocketbase/pocketbase/core"
)

func taskList(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	tasks, err := db.ListTasks()
	if err != nil {
		return err
	}

	return sendPage(e, components.TaskListPage(tasks))
}

func taskNew(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	task := models.Task{
		Id:          models.GenerateId(),
		Title:       "New Task",
		Description: "New Task Description",
	}

	if err := db.CreateTask(task); err != nil {
		return err
	}

	return e.Redirect(302, "/edit/"+task.Id)
}

func taskEdit(e *core.RequestEvent) error {
	taskId := e.Request.PathValue("id")
	if taskId == "" {
		return errors.New("task id is required")
	}

	db := database.NewDatabaseRepo(e.App)
	task, err := db.RetrieveTaskById(taskId)

	if err != nil {
		return err
	}

	return sendPage(e, components.TaskEditPage(task))
}

type TaskChange struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

func taskSave(e *core.RequestEvent) error {
	taskId := e.Request.PathValue("id")
	if taskId == "" {
		return errors.New("task id is required")
	}

	var changes TaskChange
	if err := e.BindBody(&changes); err != nil {
		return err
	}

	db := database.NewDatabaseRepo(e.App)
	task, err := db.RetrieveTaskById(taskId)
	if err != nil {
		return err
	}

	if changes.Title != "" {
		task.Title = changes.Title
	}

	if changes.Description != "" {
		task.Description = changes.Description
	}

	if err := db.UpdateTask(task); err != nil {
		return err
	}

	return e.String(200, "ok")
}
