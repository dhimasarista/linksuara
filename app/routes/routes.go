package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupRoutes(app *fiber.App, store *session.Store) {
	AuthRoutes(app)
	DashboardRoutes(app, store)
	IndexRoutes(app, store)
	InputSuaraRoutes(app, store)
	ErrorRoutes(app, store)
}
