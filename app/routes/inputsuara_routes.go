package routes

import (
	"linksuara/app/handlers"
	"linksuara/app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func InputSuaraRoutes(app *fiber.App, store *session.Store) {
	var dapil = &models.Dapil{}
	var kecamatan = &models.Kecamatan{}
	var kelurahan = &models.Kelurahan{}

	app.Get("/input-suara", func(c *fiber.Ctx) error {
		var username = handlers.GetSessionUsername(c, store) // Init username from session
		var path = c.Path()                                  // Getting Path of this routes
		//
		daftarDapil, err := dapil.FindAll()
		if err != nil {
			InternalServerError(c, err.Error())
		}
		// Render mustache(HTML) file with the data
		return c.Render("inputsuara_page", fiber.Map{
			"path":     path,
			"username": username,
			"dapil":    daftarDapil,
			"status":   fiber.StatusOK,
		})
	})

	app.Get("/input-suara/kecamatan/:id", func(c *fiber.Ctx) error {
		var idDp = c.Params("id")
		idDpNumber, _ := strconv.Atoi(idDp)

		daftarKecamatan, err := kecamatan.FindKecamatanByDapil(idDpNumber)
		if err != nil {
			handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"kecamatan": daftarKecamatan,
			"status":    fiber.StatusOK,
		})
	})

	app.Get("/input-suara/kelurahan/:id", func(c *fiber.Ctx) error {
		var idKc = c.Params("id")
		idKcNumber, _ := strconv.Atoi(idKc)

		daftarKelurahan, err := kelurahan.FindKelurahanByKecamatan(idKcNumber)
		if err != nil {
			handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"kelurahan": daftarKelurahan,
			"status":    fiber.StatusOK,
		})
	})
}
