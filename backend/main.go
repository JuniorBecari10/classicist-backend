package main

import (
	"classicist/database"
	"classicist/server"
)

func main() {
	database.CreateTables()
	server.SetupServer()
}
