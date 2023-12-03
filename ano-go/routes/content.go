package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/E8LL4IQuU/ano-go/model"
	"github.com/gofiber/fiber/v2"
)

type eventData struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
}

func getFileExtension(file *multipart.FileHeader) string {
	name := file.Filename
	ext := filepath.Ext(name)
	return ext
}

func generateSHA256Name(file io.Reader) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func GetEvents(c *fiber.Ctx) error {
	return nil
}

func CreateEvent(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "" + err.Error(),
		})
	}
	defer form.RemoveAll()

	files := form.File["image"]	// "image" is the name of the input field in the form
	if files == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "form file 'image' required",
		})
	}
	if len(files) != 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "only one image allowed",
		})
	}
	
	var path string
	for _, file := range files {
		uploadedFile, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
				"message": "" + err.Error(),
			})
		}
		defer uploadedFile.Close()

		filename, _ := generateSHA256Name(uploadedFile)
		path = filename + getFileExtension(file)
		c.SaveFile(file, "./uploads/" + path)
	}

	jsonData := form.Value["jsonData"][0]

	data := new(eventData)

	// Unmarshal JSON data into the 'data' struct
	if err := json.Unmarshal([]byte(jsonData), data); err != nil {
		return err
	}

	var event model.Event = model.Event{
		Name:			data.Name,
		Description:	data.Description,
		ImagePath:		path,
	}

	model.DB.Create(&event)

	return c.Status(fiber.StatusOK).JSON(event)
}