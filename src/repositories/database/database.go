package database

import (
	"database/sql"
	"fmt"
	"goth/src/config"
	"goth/src/models"

	_ "modernc.org/sqlite"
)

type database struct {
	conn *sql.DB
}

func NewDatabaseRepo() (*database, error) {
	conn, err := sql.Open("sqlite", config.Configs.DataBaseURL)
	if err != nil {
		return nil, err
	}

	return &database{conn}, nil
}

func (db *database) Close() error {
	return db.conn.Close()
}

func (db *database) CreateTask(task models.Task) error {
	query := `INSERT INTO tasks (id, title, description, owner_id) VALUES (?, ?, ?, ?)`
	_, err := db.conn.Exec(query, task.Id, task.Title, task.Description, task.OwnerId)
	return err
}

func (db *database) ListTasksByOwner(ownerId string) ([]models.Task, error) {
	query := `SELECT id, title, owner_id, description FROM tasks WHERE owner_id = ?`
	rows, err := db.conn.Query(query, ownerId)
	if err != nil {
		fmt.Println("query error")
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.OwnerId, &task.Description)
		if err != nil {
			fmt.Println("scan error")
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (db *database) RetrieveTaskById(taskId string) (models.Task, error) {
	query := `SELECT id, title, owner_id, description FROM tasks WHERE id = ?`
	row := db.conn.QueryRow(query, taskId)

	var task models.Task
	err := row.Scan(&task.Id, &task.Title, &task.OwnerId, &task.Description)
	if err != nil {
		return models.Task{}, err
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
