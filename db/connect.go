package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const DatabasePath = "./autoservice.db"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DatabasePath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
