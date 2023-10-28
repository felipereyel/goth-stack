package database

import (
	"database/sql"
	"goth/src/config"
	"goth/src/models"
	"strings"

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

func (db *database) UpsertUser(email string) (models.User, error) {
	user := models.User{Email: email}

	query := `SELECT id, email, name FROM users WHERE email = ?`
	row := db.conn.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Email, &user.Name)
	if err == nil {
		return user, nil
	}

	if err != sql.ErrNoRows {
		return models.EmptyUser, err
	}

	user.ID = models.GenerateId()
	user.Name = strings.Split(email, "@")[0]

	query = `INSERT INTO users (id, email, name) VALUES (?, ?, ?) RETURNING id, email, name`
	row = db.conn.QueryRow(query, user.ID, user.Email, user.Name)

	err = row.Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return models.EmptyUser, err
	}

	return user, nil
}

func (db *database) RetrieveUserById(userId string) (models.User, error) {
	query := `SELECT id, email, name FROM users WHERE id = ?`
	row := db.conn.QueryRow(query, userId)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return models.EmptyUser, err
	}

	return user, nil
}
