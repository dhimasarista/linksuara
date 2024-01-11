package routes

import (
	"linksuara/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InputSuaraRoutes(app *fiber.App, store *session.Store) {
	app.Get("/input-suara", func(c *fiber.Ctx) error {
		var path = c.Path()
		var username = handlers.GetSessionUsername(c, store)
		return c.Render("inputsuara_page", fiber.Map{
			"path":     path,
			"username": username,
		})
	})
}
