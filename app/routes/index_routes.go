package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func IndexRoutes(app *fiber.App, store *session.Store) {
	app.Get("/", func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return c.Redirect("/505")
		}
		// Mengatur session
		session.Set("username", "kotabatam")
		session.Set("user_id", "21710000")
		session.Set("logged_in", "loggedIn")
		store.CookieHTTPOnly = true
		store.CookieSecure = true
		usernameSession := session.Get("username")
		// Menyimpan Session
		if err := session.Save(); err != nil {
			log.Println(err)
			return err
		}
		log.Println("Login:", usernameSession, c.IP())

		return c.Redirect("/dashboard")
	})
}
