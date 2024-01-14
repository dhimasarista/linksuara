package routes

import (
	"linksuara/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func HitungSuaraRoutes(app *fiber.App, store *session.Store) {
	// Render Hitung Suara Page
	app.Get("/hitung-suara", func(c *fiber.Ctx) error {
		var username = handlers.GetSessionUsername(c, store) // Init username from session
		var path = c.Path()                                  // Getting Path of this routes
		// Render mustache(HTML) file with the data
		return c.Render("hitungsuara_page", fiber.Map{
			"username": username,
			"path":     path,
		})
	})
}
