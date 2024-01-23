package handlers

import (
	"linksuara/app/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	buktiSuara := &models.BuktiSuara{}
	form, err := c.MultipartForm() // Init Multipartform
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	var filename string
	// Mengambil files dengan key image dari map
	files := form.File["image"]
	for _, file := range files {
		err := c.SaveFile(file, "./app/uploads/images/"+file.Filename)
		filename = file.Filename
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		// err = CompressImage("./app/uploads/images/"+file.Filename, "./app/uploads/compressed/"+file.Filename, 200)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return c.JSON(fiber.Map{
		// 		"status": fiber.StatusBadRequest,
		// 		"error":  err.Error(),
		// 	})
		// }

		// err = os.Remove("./app/uploads/images" + file.Filename)
		// if err != nil {
		// 	return c.JSON(fiber.Map{
		// 		"status": fiber.StatusBadRequest,
		// 		"error":  err.Error(),
		// 	})
		// }
	}

	// Mengiri Foto ke MySQL
	err = buktiSuara.NewData(filename, "New")
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  c.Response().StatusCode(),
		"message": "Image uploaded successfully",
	})
}

func DeleteImage(c *fiber.Ctx) error {
	// filename := c.Params("filename")
	var formData map[string]string
	err := c.BodyParser(&formData)
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	// Menghapus file
	err = os.Remove("./app/uploads/images/" + string(formData["filename"]))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	// Mengirim response 200 ke client
	return c.JSON(fiber.Map{
		"status":  c.Response().StatusCode(),
		"message": "Success delete image",
	})
}

// func CompressImage(inputPath, outputPath string, maxFileSize int) error {
// 	file, err := os.Open(inputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return err
// 	}

// 	// Resize the image to achieve the target file size
// 	// You may need to experiment with the resize dimensions to achieve the desired file size
// 	resizedImg := resize.Resize(0, 0, img, resize.Lanczos3)

// 	// Create a new output file
// 	outputFile, err := os.Create(outputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()

// 	// Encode the resized image to JPEG format
// 	err = jpeg.Encode(outputFile, resizedImg, nil)
// 	if err != nil {
// 		return err
// 	}

// 	// Check if the file size is below the maximum limit
// 	fileInfo, err := outputFile.Stat()
// 	if err != nil {
// 		return err
// 	}

// 	if int(fileInfo.Size()) > maxFileSize {
// 		return fmt.Errorf("failed to compress image to the target file size")
// 	}

// 	return nil
// }
