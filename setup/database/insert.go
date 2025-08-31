package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertBaseData() {
	log.Println("Inserting base data into the database..")

	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	inserts := `
-- TODO!
`
	_, err = db.Exec(inserts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Base data insertion complete.")
}