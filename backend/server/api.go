package server

import (
	"classicist/database"
	"classicist/query"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const API_GROUP = "/api"

func setupApi(app *fiber.App) {
	api := app.Group(API_GROUP)

	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Start")
	})

	getById(api, "composer", "Composer", query.GetComposerById)
	getById(api, "work", "Work", query.GetWorkById)
}

func getById[W any](api fiber.Router, route, what string, getFn func (*sql.DB, int) (W, error)) fiber.Router {
	return api.Get(fmt.Sprintf("/%s/:id", route), func (c *fiber.Ctx) error {
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

		it, err := getFn(db, id)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("%s not found", what),
			})
		}

		return c.JSON(it)
	})
}
