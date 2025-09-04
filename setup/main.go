package main

import (
	"log"
	"setup/database"
	"setup/public"
)

func main() {
	log.Println("Starting setup..")

	database.CreateTables()
	database.InsertSeed()
	public.FetchPublicFolder()

	log.Println("Setup complete. You may start the server safely now.")
}
