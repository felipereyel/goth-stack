package controllers

import (
	"goth/src/models"
	"goth/src/repositories/database"

	"github.com/google/uuid"
)

type TaskController struct {
	DbRepo database.Database
}

func NewTaskController(dbRepo database.Database) *TaskController {
	return &TaskController{dbRepo}
}

// PROTECTED

func (tc *TaskController) CreateTask(ownerId string) (models.Task, error) {
	task := models.Task{
		Id:          uuid.New().String(),
		Title:       "New Task",
		Description: "New Task Description",
		OwnerId:     ownerId,
	}
	return task, tc.DbRepo.CreateTask(task)
}

func (tc *TaskController) ListTasks(ownerId string) ([]models.Task, error) {
	return tc.DbRepo.ListTasksByOwner(ownerId)
}

// SCOPED

func (tc *TaskController) RetrieveTask(taskId string) (models.Task, error) {
	return tc.DbRepo.RetrieveTaskById(taskId)
}

func (tc *TaskController) DeleteTask(taskId string) error {
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
