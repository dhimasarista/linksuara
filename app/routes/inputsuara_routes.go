package routes

import (
	"linksuara/app/handlers"
	"linksuara/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InputSuaraRoutes(app *fiber.App, store *session.Store) {
	var dapil = &models.Dapil{}
	var kecamatan = &models.Kecamatan{}
	var kelurahan = &models.Kelurahan{}

	app.Get("/input-suara", func(c *fiber.Ctx) error {
		var path = c.Path()
		var username = handlers.GetSessionUsername(c, store)

		daftarDapil, err := dapil.FindAll()
		if err != nil {
			InternalServerError(c, err.Error())
		}
		daftarKecamatan, err := kecamatan.FindAll()
		if err != nil {
			InternalServerError(c, err.Error())
		}

		daftarKelurahan, err := kelurahan.FindAll()
		if err != nil {
			InternalServerError(c, err.Error())
		}

		return c.Render("inputsuara_page", fiber.Map{
			"path":      path,
			"username":  username,
			"dapil":     daftarDapil,
			"kecamatan": daftarKecamatan,
			"kelurahan": daftarKelurahan,
		})
	})
}
