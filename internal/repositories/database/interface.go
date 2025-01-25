package database

import (
	"goth/internal/models"

	_ "modernc.org/sqlite"
)

type Database interface {
	Close() error

	CreateTask(task models.Task) error
	RetrieveTaskById(taskId string) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(taskId string) error
	ListTasksByOwner(ownerId string) ([]models.Task, error)

	UpsertUser(email string) (models.User, error)
	RetrieveUserById(userId string) (models.User, error)
}
