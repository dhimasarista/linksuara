package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupRoutes(app *fiber.App, store *session.Store) {
	AuthRoutes(app)
	IndexRoutes(app, store)
	DashboardRoutes(app, store)
	InputSuaraRoutes(app, store)
	HitungSuaraRoutes(app, store)
	ErrorRoutes(app, store)
}
