package controllers

import (
	"database/sql"
	"goth/internal/models"
	"goth/internal/repositories/database"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	DbRepo database.Database
}

func NewTaskController(dbRepo database.Database) *TaskController {
	return &TaskController{
		DbRepo: dbRepo,
	}
}

func (tc *TaskController) CreateTask() (models.Task, error) {
	task := models.Task{
		Id:          models.GenerateId(),
		Title:       "New Task",
		Description: "New Task Description",
	}
	return task, tc.DbRepo.CreateTask(task)
}

func (tc *TaskController) ListTasks() ([]models.Task, error) {
	return tc.DbRepo.ListTasks()
}

func (tc *TaskController) RetrieveTask(taskId string) (models.Task, error) {
	task, err := tc.DbRepo.RetrieveTaskById(taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.EmptyTask, fiber.ErrNotFound
		}

		return models.EmptyTask, err
	}

	return task, nil
}

func (tc *TaskController) DeleteTask(taskId string) error {
	_, err := tc.RetrieveTask(taskId)
	if err != nil {
		return err
	}

	return tc.DbRepo.DeleteTask(taskId)
}

type TaskChange struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (tc *TaskController) UpdateTask(taskId string, changes TaskChange) error {
	task, err := tc.RetrieveTask(taskId)
	if err != nil {
		return err
	}

	if changes.Title != "" {
		task.Title = changes.Title
	}

	if changes.Description != "" {
		task.Description = changes.Description
	}

	return tc.DbRepo.UpdateTask(task)
}
