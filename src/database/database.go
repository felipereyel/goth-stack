package database

import (
	"database/sql"
	"goth/src/config"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

func (db *Database) Close() error {
	return db.Conn.Close()
}

func New() (*Database, error) {
	conn, err := sql.Open("sqlite3", config.Configs.DataBaseURL)
	if err != nil {
		return nil, err
	}

	return &Database{Conn: conn}, nil
}
