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
	query := `INSERT INTO tasks (id, title, description, owner_id) VALUES (?, ?, ?, ?)`
	_, err := db.conn.Exec(query, task.Id, task.Title, task.Description, task.OwnerId)
	return err
}

func (db *database) ListTasksByOwner(ownerId string) ([]models.Task, error) {
	query := `SELECT id, title, owner_id, description FROM tasks WHERE owner_id = ?`
	rows, err := db.conn.Query(query, ownerId)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.OwnerId, &task.Description)
		if err != nil {
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

func (db *database) InsertUser(user models.User) error {
	query := `INSERT INTO users (id, username, pswd_hash) VALUES (?, ?, ?)`
	_, err := db.conn.Exec(query, user.ID, user.Username, user.PswdHash)
	return err
}

func (db *database) RetrieveUserByName(username string) (models.User, error) {
	query := `SELECT id, username, pswd_hash FROM users WHERE username = ?`
	row := db.conn.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.PswdHash)
	if err != nil {
		return models.EmptyUser, err
	}

	return user, nil
}

func (db *database) RetrieveUserById(id string) (models.User, error) {
	query := `SELECT id, username, pswd_hash FROM users WHERE id = ?`
	row := db.conn.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.PswdHash)
	if err != nil {
		return models.EmptyUser, err
	}

	return user, nil
}
