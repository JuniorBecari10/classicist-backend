package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const PORT = 8080

func SetupServer() {
	app := fiber.New()
	setupApi(app)
	
	app.Listen(fmt.Sprintf(":%d", PORT))
}
