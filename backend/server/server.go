package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const PORT = 8080

func SetupServer() {
	log.Printf("Starting server on port %d..\n", PORT)
	app := fiber.New()

	app.Static("/public", "../public")
	setupApi(app)
	
	app.Listen(fmt.Sprintf(":%d", PORT))
}
