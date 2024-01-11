package routes

import (
	"linksuara/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func HitungSuaraRoutes(app *fiber.App, store *session.Store) {
	app.Get("/hitung-suara", func(c *fiber.Ctx) error {
		var username = handlers.GetSessionUsername(c, store)
		var path = c.Path()
		return c.Render("hitungsuara_page", fiber.Map{
			"username": username,
			"path":     path,
		})
	})
}