package database

import (
	"database/sql"
	"goth/internal/config"
	"goth/internal/models"

	_ "modernc.org/sqlite"
)

type database struct {
	conn *sql.DB
}

func NewDatabaseRepo(cfg config.ServerConfigs) (Database, error) {
	conn, err := sql.Open("sqlite", cfg.DataBaseURL)
	if err != nil {
		return nil, err
	}

	return &database{conn}, nil
}

func (db *database) Close() error {
	return db.conn.Close()
}

func (db *database) CreateTask(task models.Task) error {
	query := `INSERT INTO tasks (id, title, description) VALUES (?, ?, ?)`
	_, err := db.conn.Exec(query, task.Id, task.Title, task.Description)
	return err
}

func (db *database) ListTasks() ([]models.Task, error) {
	query := `SELECT id, title, description FROM tasks`
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (db *database) RetrieveTaskById(taskId string) (models.Task, error) {
	query := `SELECT id, title, description FROM tasks WHERE id = ?`
	row := db.conn.QueryRow(query, taskId)

	var task models.Task
	err := row.Scan(&task.Id, &task.Title, &task.Description)
	if err != nil {
		return models.EmptyTask, err
	}

	return task, nil
}

func (db *database) DeleteTask(taskId string) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := db.conn.Exec(query, taskId)
	return err
}

func (db *database) UpdateTask(task models.Task) error {
	query := `UPDATE tasks SET title = ?, description = ? WHERE id = ?`
	_, err := db.conn.Exec(query, task.Title, task.Description, task.Id)
	return err
}
