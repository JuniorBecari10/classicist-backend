package server

import (
	"classicist/database"
	"classicist/query"

	"github.com/gofiber/fiber/v2"
)

const API_GROUP = "/api"

func setupApi(app *fiber.App) {
	api := app.Group(API_GROUP)

	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Start")
	})

	api.Get("/composer/:id", func (c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		db, err := database.GetDatabaseConnection()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot connect to database",
			})
		}

		composer, err := query.GetComposerById(db, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Composer not found",
			})
		}

		return c.JSON(composer)
	})
}
