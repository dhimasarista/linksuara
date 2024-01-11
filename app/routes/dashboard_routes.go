package routes

import (
	"linksuara/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var username = handlers.GetSessionUsername(c, store)
		var path = c.Path()

		return c.Render("dashboard_page", fiber.Map{
			"username": username,
			"path":     path,
		})
	})
}
