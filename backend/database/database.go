package database

import (
	"database/sql"
	"shared/dbpath"
)

const DATABASE = "database.db"

func GetDatabasePath() string {
	return dbpath.GetPath(DATABASE)
}

// Don't forget to defer the closing of the handler.
func GetDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", GetDatabasePath())
	return db, err
}
