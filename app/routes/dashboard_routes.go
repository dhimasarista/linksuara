package routes

import (
	"linksuara/app/handlers"
	"linksuara/app/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DashboardRoutes(app *fiber.App, store *session.Store) {
	dapil := &models.Dapil{}
	kecamatan := &models.Kecamatan{}
	kelurahan := &models.Kelurahan{}
	tps := &models.TPS{}
	// Render Dashboard Page
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var username = handlers.GetSessionUsername(c, store) // Init username from session
		var path = c.Path()                                  // Getting Path of this routes
		// Counting total of dapil
		totalDapil, err := dapil.TotalDapil()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		// Counting total of kecamatan
		totalKecamatan, err := kecamatan.TotalKecamatan()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		// Counting total of kelurahan
		totalKelurahan, err := kelurahan.TotalKelurahan()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		// Counting total of TPS
		totalTPS, err := tps.TotalTPS()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		// Render mustache(HTML) file with the data
		return c.Render("dashboard_page", fiber.Map{
			"username":        username,
			"path":            path,
			"total_dapil":     totalDapil,
			"total_kecamatan": totalKecamatan,
			"total_kelurahan": totalKelurahan,
			"total_tps":       totalTPS,
			"status":          fiber.StatusOK,
		})
	})
}
