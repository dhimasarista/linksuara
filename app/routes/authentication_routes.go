package routes

import "github.com/gofiber/fiber/v2"

func AuthRoutes(app *fiber.App) {
	// Render Login Page
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login_page", fiber.Map{})
	})
}
