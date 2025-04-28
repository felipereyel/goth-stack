package database

import (
	"goth/internal/models"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type database struct {
	app core.App
}

func NewDatabaseRepo(app core.App) database {
	return database{app}
}

func (db *database) CreateTask(task models.Task) error {
	query := `INSERT INTO tasks (id, title, description) VALUES ({:Id}, {:Title}, {:Description})`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id":          task.Id,
		"Title":       task.Title,
		"Description": task.Description,
	})

	_, err := q.Execute()
	return err
}

func (db *database) ListTasks() ([]models.Task, error) {
	query := `SELECT id, title, description FROM tasks`
	q := db.app.DB().NewQuery(query)

	tasks := make([]models.Task, 0)
	if err := q.All(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *database) RetrieveTaskById(taskId string) (models.Task, error) {
	query := `SELECT id, title, description FROM tasks WHERE id = {:Id}`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id": taskId,
	})

	var task models.Task
	err := q.One(&task)
	return task, err
}

func (db *database) DeleteTask(taskId string) error {
	query := `DELETE FROM tasks WHERE id = {:Id}`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id": taskId,
	})

	_, err := q.Execute()
	return err
}

func (db *database) UpdateTask(task models.Task) error {
	query := `UPDATE tasks SET title = {:Title}, description = {:Description} WHERE id = {:Id}`
	q := db.app.DB().NewQuery(query).Bind(dbx.Params{
		"Id":          task.Id,
		"Title":       task.Title,
		"Description": task.Description,
	})

	_, err := q.Execute()
	return err
}
