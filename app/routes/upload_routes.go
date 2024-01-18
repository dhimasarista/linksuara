package routes

import (
	"linksuara/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func FileManagement(app *fiber.App) {
	app.Post("/upload/image", handlers.UploadImage)
	app.Delete("/delete/image", handlers.DeleteImage)
}
