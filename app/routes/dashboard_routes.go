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
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		var username = handlers.GetSessionUsername(c, store)
		var path = c.Path()

		totalDapil, err := dapil.TotalDapil()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		totalKecamatan, err := kecamatan.TotalKecamatan()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}
		totalKelurahan, err := kelurahan.TotalKelurahan()
		if err != nil {
			log.Println(err)
			return InternalServerError(c, err.Error())
		}

		return c.Render("dashboard_page", fiber.Map{
			"username":        username,
			"path":            path,
			"total_dapil":     totalDapil,
			"total_kecamatan": totalKecamatan,
			"total_kelurahan": totalKelurahan,
			"status":          fiber.StatusOK,
		})
	})
}
