package handlers

import "github.com/gofiber/fiber/v2"

// Response in JSON form
func ResponseJSON(c *fiber.Ctx, message string, status int) error {
	return c.JSON(fiber.Map{
		"message": message,
		"status":  status,
	})
}
