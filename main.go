package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engine := mustache.New("./views", ".mustache") // Init template engine
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/404")
		},
	})

	app.Use(recover.New())      // Init recover with default config
	app.Static("/", "./public") // Middleware untuk menyediakan file static

	app.Get("/metrics", monitor.New())

	log.Fatal(app.Listen(":9999"))
}
