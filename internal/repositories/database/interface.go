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

	InsertUser(user models.User) error
	RetrieveUserById(id string) (models.User, error)
	RetrieveUserByName(username string) (models.User, error)
}
