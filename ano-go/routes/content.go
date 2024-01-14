package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/E8LL4IQuU/ano-go/model"
	"github.com/gofiber/fiber/v2"
)

type eventData struct {
	Title		string	`json:"title"`
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

func saveImage(c *fiber.Ctx, isImageRequired bool) (*multipart.Form, string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "" + err.Error(),
		})
	}
	defer form.RemoveAll()


	files := form.File["image"]	// "image" is the name of the input field in the form

	if isImageRequired {
		if files == nil {
			return nil, "", c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "form file 'image' required",
			})
		}
	} else {
		if files == nil {
			return nil, "", nil
		}
	}


	if len(files) != 1 {
		return nil, "", c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "only one image allowed",
		})
	}
	
	var path string
	for _, file := range files {
		uploadedFile, err := file.Open()
		if err != nil {
			return nil, "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
				"message": "" + err.Error(),
			})
		}
		defer uploadedFile.Close()

		filename, _ := generateSHA256Name(uploadedFile)
		path = filename + getFileExtension(file)
		c.SaveFile(file, "./uploads/" + path)
	}

	return form, path, nil
}

func GetEvents(c *fiber.Ctx) error {
	// TODO: control limit by c.Params() I guess
	var events []model.Event
	model.DB.Order("ID desc").Limit(10).Find(&events)

	return c.Status(fiber.StatusOK).JSON(events)
}

func CreateEvent(c *fiber.Ctx) error {

	form, path, err := saveImage(c, true)
	if err != nil {
		return err
	}

	jsonData := form.Value["jsonData"][0]

	data := new(eventData)

	// Unmarshal JSON data into the 'data' struct
	if err := json.Unmarshal([]byte(jsonData), data); err != nil {
		return err
	}

	var event model.Event = model.Event{
		Title:			data.Title,
		Description:	data.Description,
		ImagePath:		path,
	}

	model.DB.Create(&event)

	return c.Status(fiber.StatusOK).JSON(event)
}

func UpdateEvent(c *fiber.Ctx) error {
	// Parse event ID from the request params
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Find the event by ID in the database
	var event model.Event
	if err := model.DB.First(&event, id).Error; err != nil {
		// Handle the error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Event not found"})
	}

	// Parse the event details
	var updatedEvent model.Event
	if err := c.BodyParser(&updatedEvent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Parse the image
	_, updatedPath, err := saveImage(c, false)
	if err != nil {
		return err
	}

	// Update the book details
	if updatedEvent.Title != "" {
		event.Title = updatedEvent.Title
	}
	if updatedEvent.Description != "" {
		event.Description = updatedEvent.Description
	}
	if updatedPath != "" {
		event.ImagePath = updatedPath
	}

	// Save the updated event back to the database
	model.DB.Save(&event)

	return c.Status(fiber.StatusOK).JSON(event)
}

func DeleteEvent(c *fiber.Ctx) error {

	eventID := c.Params("id")

	// Find record in the database
	var event model.Event
	if err := model.DB.First(&event, eventID).Error; err != nil {
		// Handle the error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	// Delete the record
	if err := model.DB.Delete(&event).Error; err != nil {
		// Handle the error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting record"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Record deleted successfully"})
}