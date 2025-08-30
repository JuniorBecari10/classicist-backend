package main

import (
	"classicist/database"
	"classicist/server"
)

func main() {
	setupEnvironment()
	server.SetupServer()
}

func setupEnvironment() {
	database.CreateTables()
}
