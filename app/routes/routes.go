package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Setup All Routes
func SetupRoutes(app *fiber.App, store *session.Store) {
	AuthRoutes(app)
	ErrorRoutes(app, store)
	IndexRoutes(app, store)
	DashboardRoutes(app, store)
	InputSuaraRoutes(app, store)
	HitungSuaraRoutes(app, store)
	FileManagement(app)
}
