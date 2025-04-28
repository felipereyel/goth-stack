package database

import (
	"database/sql"
	"goth/internal/models"

	_ "modernc.org/sqlite"
)

type fakeDatabase struct {
	tasks map[string]models.Task
}

func NewFakeDatabaseRepo() (Database, error) {
	return &fakeDatabase{
		tasks: make(map[string]models.Task),
	}, nil
}

func (db *fakeDatabase) Close() error {
	return nil
}

func (db *fakeDatabase) CreateTask(task models.Task) error {
	db.tasks[task.Id] = task
	return nil
}

func (db *fakeDatabase) ListTasks() ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	for _, task := range db.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (db *fakeDatabase) RetrieveTaskById(taskId string) (models.Task, error) {
	task, ok := db.tasks[taskId]
	if !ok {
		return models.EmptyTask, sql.ErrNoRows
	}

	return task, nil
}

func (db *fakeDatabase) DeleteTask(taskId string) error {
	delete(db.tasks, taskId)
	return nil
}

func (db *fakeDatabase) UpdateTask(task models.Task) error {
	db.tasks[task.Id] = task
	return nil
}
