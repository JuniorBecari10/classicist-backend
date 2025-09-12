package database

import (
	"database/sql"
	"log"
	"os"
	"shared/dbpath"

	_ "github.com/mattn/go-sqlite3"
	_ "embed"
)

//go:embed sql/schema.sql
var schema string

const FILE = "database.db"

func CreateTables() {
	log.Println("Erasing previous database file if it exists..")

	path := dbpath.GetPath(FILE)

	err := os.Remove(path)
	if err != nil {
		log.Println("Cannot delete:", err)
		log.Println("Continuing..")
	}

	log.Println("Creating database schema..")

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database schema created.")
}
