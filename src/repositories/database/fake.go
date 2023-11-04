package database

import (
	"database/sql"
	"goth/src/models"

	_ "modernc.org/sqlite"
)

type fakeDatabase struct {
	tasks map[string]models.Task
	users map[string]models.User
}

func NewFakeDatabaseRepo() (Database, error) {
	return &fakeDatabase{
		tasks: make(map[string]models.Task),
		users: make(map[string]models.User),
	}, nil
}

func (db *fakeDatabase) Close() error {
	return nil
}

func (db *fakeDatabase) CreateTask(task models.Task) error {
	db.tasks[task.Id] = task
	return nil
}

func (db *fakeDatabase) ListTasksByOwner(ownerId string) ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	for _, task := range db.tasks {
		if task.OwnerId == ownerId {
			tasks = append(tasks, task)
		}
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

func (db *fakeDatabase) UpsertUser(email string) (models.User, error) {
	for _, u := range db.users {
		if u.Email == email {
			return u, nil
		}
	}

	user := models.User{
		Email: email,
		ID:    models.GenerateId(),
		Name:  models.GenerateNameFromEmail(email),
	}

	db.users[user.ID] = user

	return user, nil
}

func (db *fakeDatabase) RetrieveUserById(userId string) (models.User, error) {
	user, ok := db.users[userId]
	if !ok {
		return models.EmptyUser, sql.ErrNoRows
	}

	return user, nil
}
