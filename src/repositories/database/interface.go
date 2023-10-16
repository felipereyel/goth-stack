package database

import (
	"goth/src/models"

	_ "modernc.org/sqlite"
)

// CRUDL
type Database interface {
	Close() error
	CreateTask(task models.Task) error
	RetrieveTaskById(taskId string) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(taskId string) error
	ListTasksByOwner(ownerId string) ([]models.Task, error)
}
