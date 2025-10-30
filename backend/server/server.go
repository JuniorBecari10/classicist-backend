package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const PORT = 8080

func SetupServer() {
	log.Printf("Starting server on port %d..\n", PORT)
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/public", "../public/files")
	setupApi(app)
	
	app.Listen(fmt.Sprintf(":%d", PORT))
}
