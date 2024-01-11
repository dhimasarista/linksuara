package routes

import "github.com/gofiber/fiber/v2"

func InputSuaraRoutes(app *fiber.App) {

	app.Get("/input-suara", func(c *fiber.Ctx) error {
		return c.Render("inputsuara_page", fiber.Map{})
	})
}
