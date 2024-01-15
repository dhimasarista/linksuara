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
	var tps = &models.TPS{}
	var caleg = &models.Caleg{}
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

	app.Get("/pilih/kecamatan/:id", func(c *fiber.Ctx) error {
		var idDp = c.Params("id")
		idDpNumber, _ := strconv.Atoi(idDp)

		daftarKecamatan, err := kecamatan.FindKecamatanByDapil(idDpNumber)
		if err != nil {
			return handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"kecamatan": daftarKecamatan,
			"status":    fiber.StatusOK,
		})
	})

	app.Get("/pilih/kelurahan/:id", func(c *fiber.Ctx) error {
		var idKc = c.Params("id")
		idKcNumber, _ := strconv.Atoi(idKc)

		daftarKelurahan, err := kelurahan.FindKelurahanByKecamatan(idKcNumber)
		if err != nil {
			return handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"kelurahan": daftarKelurahan,
			"status":    fiber.StatusOK,
		})
	})

	app.Get("/pilih/caleg/:id", func(c *fiber.Ctx) error {
		var idDapil = c.Params("id")
		idDapilNumber, _ := strconv.Atoi(idDapil)

		dataCaleg, err := caleg.FindAllByDapil(idDapilNumber)
		if err != nil {
			return handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"caleg":  dataCaleg,
			"status": fiber.StatusOK,
		})
	})

	app.Get("/pilih/tps/:id", func(c *fiber.Ctx) error {
		var idKl = c.Params("id")
		idKlNumber, _ := strconv.Atoi(idKl)

		daftarTps, err := tps.FindTpsByKelurahan(idKlNumber)
		if err != nil {
			return handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"tps":    daftarTps,
			"status": fiber.StatusOK,
		})
	})

	app.Get("/input-suara/tps/:id", func(c *fiber.Ctx) error {
		var idKl = c.Params("id")
		idKlNumber, _ := strconv.Atoi(idKl)

		dataTps, err := tps.GetByID(idKlNumber)
		if err != nil {
			return handlers.ResponseJSON(c, err.Error(), fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"data":   dataTps,
			"status": fiber.StatusOK,
		})
	})
}
