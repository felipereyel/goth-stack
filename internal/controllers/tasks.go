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

func (tc *TaskController) CreateTask(ownerId string) (models.Task, error) {
	task := models.Task{
		Id:          models.GenerateId(),
		Title:       "New Task",
		Description: "New Task Description",
		OwnerId:     ownerId,
	}
	return task, tc.DbRepo.CreateTask(task)
}

func (tc *TaskController) ListTasks(ownerId string) ([]models.Task, error) {
	return tc.DbRepo.ListTasksByOwner(ownerId)
}

func (tc *TaskController) RetrieveTask(ownerId, taskId string) (models.Task, error) {
	task, err := tc.DbRepo.RetrieveTaskById(taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.EmptyTask, fiber.ErrNotFound
		}

		return models.EmptyTask, err
	}

	if task.OwnerId != ownerId {
		return models.EmptyTask, fiber.ErrNotFound
	}

	return task, nil
}

func (tc *TaskController) DeleteTask(ownerId, taskId string) error {
	_, err := tc.RetrieveTask(ownerId, taskId)
	if err != nil {
		return err
	}

	return tc.DbRepo.DeleteTask(taskId)
}

type TaskChange struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (tc *TaskController) UpdateTask(ownerId, taskId string, changes TaskChange) error {
	task, err := tc.RetrieveTask(ownerId, taskId)
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
