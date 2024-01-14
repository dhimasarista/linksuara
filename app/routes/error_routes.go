package routes

import (
	"fmt"
	"linksuara/app/handlers"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Routes of Error Handling
func ErrorRoutes(app *fiber.App, store *session.Store) {
	// Render Error Page
	app.Get("/error", func(c *fiber.Ctx) error {
		var path string = c.Path()
		var username string = handlers.GetSessionUsername(c, store)
		// Render mustache(HTML) file with the data
		return c.Render("error_page", fiber.Map{
			"path": path,
			"user": username,
		})
	})
	// Path of error will redirect with query paramater, specific http code and message of error
	// Client-side will render http code and message from query parameter
	app.Get("/500", func(c *fiber.Ctx) error {
		// Internal Server Error with Default Message
		return c.Redirect("/error?code=500&title=Internal+Server+Error&message=We+will+fix+it+as+soon+as+possible...")
	})
	app.Get("/404", func(c *fiber.Ctx) error {
		// Error Not Found
		return c.Redirect("/error?code=404&title=Page+Not+Found&message=It+looks+like+you+found+a+glitch+in+the+matrix...")
	})
	app.Get("/418", func(c *fiber.Ctx) error {
		// Joking Error with Tea Pot
		return c.Redirect("/error?code=418&title=I`am+a+Teapot&message=hahahahahahhahahahahahahhahaha...")
	})
}

// Internal Server Errir with Specific Error from Server
func InternalServerError(c *fiber.Ctx, message string) error {
	log.Println(message)                                          // Show the Error to Console
	var messageFormatted = strings.Replace(message, " ", "+", -1) // Formatted message for fit the query parameter
	var path = fmt.Sprintf("/error?code=500&title=Internal+Server+Error&message=%s", messageFormatted)
	return c.Redirect(path) // Redirect Path to Error Page
}

// func HandleErrorAndRollback(tx *gorm.Tx, err error, c *fiber.Ctx) error {
// 	if err != nil {
// 		log.Println(err)
// 		tx.Rollback()
// 		return c.JSON(fiber.Map{
// 			"error":  err.Error(),
// 			"status": fiber.StatusInternalServerError,
// 		})
// 	}
// 	return nil
// }
