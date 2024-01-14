package main

import (
	"linksuara/app/routes"
	"linksuara/app/utility"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/mustache/v2"
)

var store = *session.New(session.ConfigDefault) // Storing session

func main() {
	utility.ClearScreen() // Clearing console

	engine := mustache.New("./views", ".mustache") // Init template engine
	app := fiber.New(fiber.Config{
		Views: engine, // Using template engine
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/404") // Handling for nothing routes
		},
	})

	app.Use(recover.New())      // Init recover with default config
	app.Static("/", "./public") // Middleware untuk menyediakan file static

	// Routes
	routes.SetupRoutes(app, &store)    // All routes is Here
	app.Get("/metrics", monitor.New()) // Show metrics

	log.Fatal(app.Listen(":9999")) // Show the console of app
}
