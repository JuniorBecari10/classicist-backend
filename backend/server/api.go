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

	apiEndpoint(api, "composer", "Composer", query.QueryComposer)
	apiEndpoint(api, "work", "Work", query.QueryWork)
	apiEndpoint(api, "performer", "Performer", query.QueryPerformer)
	apiEndpoint(api, "recsforwork", "Recordings", query.QueryRecordingsForWork)
	apiEndpoint(api, "recsbyperformer", "Recordings", query.QueryRecordingsByPerformer)
	apiEndpoint(api, "worksbycomposer", "Works", query.QueryWorksByComposer)

	api.Get("/search", searchEndpoint)
}

func apiEndpoint[W any](api fiber.Router, route, what string, query func (*sql.DB, int) (W, error)) fiber.Router {
	return api.Get(fmt.Sprintf("/%s", route), func (c *fiber.Ctx) error {
		id := c.QueryInt("id", -1)
		if id < 0 {
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
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("%s not found", what),
			})
		}

		return c.JSON(it)
	})
}

func searchEndpoint(c *fiber.Ctx) error {
	term := c.Query("q")

	db, err := database.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect to database",
		})
	}

	results, err := query.SearchDatabase(db, term)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cannot get search results",
		})
	}

	return c.JSON(results)
}
