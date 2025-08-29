package server

import "github.com/gofiber/fiber/v2"

func setupApi(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Start")
	})
	
	api.Get("/hello", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
