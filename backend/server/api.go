package server

import "github.com/gofiber/fiber/v2"

const API_GROUP = "/api"

func setupApi(app *fiber.App) {
	api := app.Group(API_GROUP)

	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Start")
	})
	
	api.Get("/hello", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
