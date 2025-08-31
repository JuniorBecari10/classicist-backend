package main

import (
	"classicist/server"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	checkRequirements()
	server.SetupServer()
}

func checkRequirements() {
	log.Println("Checking system requirements..")

	checkDatabase()
	checkPublicFolder()

	log.Println("All requirements met. The server may start safely.")
}

func checkDatabase() {
	log.Println("Checking database connection..")

	// Check if 'database.db' exists and is usable via ping.

	// Check if the the file exists.
	if _, err := os.Stat("../database.db"); os.IsNotExist(err) {
		log.Fatal("ERROR: Database file does not exist.")
	}

	// Check if a connection is possible.
	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		log.Fatal("ERROR:", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("ERROR:", err)
	}

	log.Println("Database connection successful.")
}

func checkPublicFolder() {
	log.Println("Checking public folder..")

	// Check if the public folder exists.
	if _, err := os.Stat("../public"); os.IsNotExist(err) {
		log.Fatal("ERROR: Public folder does not exist.")
	}

	log.Println("Public folder exists.")
}
