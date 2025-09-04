package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "embed"
)

//go:embed sql/seed.sql
var seed string

func InsertSeed() {
	log.Println("Inserting base data into the database..")

	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(seed)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Base data insertion complete.")
}
