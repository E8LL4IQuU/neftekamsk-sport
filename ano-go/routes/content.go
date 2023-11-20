package routes

import (
	"github.com/gofiber/fiber/v2"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func generateSHA256Name(file io.Reader) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "" + err.Error(),
		})
	}

	files := form.File["file"]	// "file" is the name of the input field in the form

	for _, file := range files {
		uploadedFile, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
				"message": "" + err.Error(),
			})
		}
		defer uploadedFile.Close()

		filename, err := generateSHA256Name(uploadedFile)

		c.SaveFile(file, "./uploads/"+filename)
	}

	return nil
}