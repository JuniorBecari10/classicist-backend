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

	apiEndpoint(api, "composer", "Composer", query.GetComposerById)
	apiEndpoint(api, "work", "Work", query.GetWorkById)
	apiEndpoint(api, "recbywork", "Recording", query.GetRecordingsByWorkId)
}

func apiEndpoint[W any](api fiber.Router, route, what string, query func (*sql.DB, int) (W, error)) fiber.Router {
	return api.Get(fmt.Sprintf("/%s/:id", route), func (c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		db, err := database.GetDatabaseConnection()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot connect to database",
			})
		}

		it, err := query(db, id)
		log.Println(err)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("%s not found", what),
			})
		}

		return c.JSON(it)
	})
}
