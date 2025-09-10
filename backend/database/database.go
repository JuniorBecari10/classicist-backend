package database

import "database/sql"

const DATABASE_PATH = "../database.db"

// Don't forget to defer the closing of the handler.
func GetDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DATABASE_PATH)
	return db, err
}
