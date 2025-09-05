package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertSeed() {
	log.Println("Inserting seed data into the database..")

	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var seed string

	_, err = db.Exec(seed)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Seed data insertion complete.")
}
