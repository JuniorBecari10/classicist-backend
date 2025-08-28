package main

import (
	"classicist/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.CreateTables()
	setupServer()
}

func setupServer() {
	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Start")
	})
	
	app.Get("/hello", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	app.Listen(":8080")
}
