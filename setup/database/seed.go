package database

import (
	"database/sql"
	"log"
	"shared/serialization"

	_ "github.com/mattn/go-sqlite3"
)

func InsertSeed() {
	log.Println("Inserting seed data into the database..")

	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	seed := serialization.WriteAll(works, composers)

	_, err = db.Exec(seed)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Seed data insertion complete.")
}
