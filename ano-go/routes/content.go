package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/E8LL4IQuU/ano-go/model"
	"github.com/gofiber/fiber/v2"
)

// TODO: We'll have to send random data to this server to check every possible panic it gives

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

func getItems(c *fiber.Ctx, modelType interface{}, orderByColumn string) error {
	limit := 10

	// Check if a custom limit is provided in the request parameters
	if limitParam := c.Params("limit"); limitParam != "" {
		// Attempt to parse the limit parameter to an integer
		if customLimit, err := strconv.Atoi(limitParam); err == nil && customLimit > 0 {
			limit = customLimit
		}
	}

	// Create a new instance of the modelType
	item := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()

	// Retrieve the value of the created instance
	itemValue := reflect.ValueOf(item)

	if err := model.DB.Order(orderByColumn + " desc").Limit(limit).Find(itemValue.Interface()).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(itemValue.Interface())
}

func createItem(c *fiber.Ctx, isImageRequired bool, modelType interface{}) error {
	// FIXME: crashes when sending an empty json
	form, path, err := saveImage(c, isImageRequired)
	if err != nil {
		return err
	}

	title := form.Value["title"][0]
	description := form.Value["description"][0]

	item := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()
	reflect.ValueOf(item).Elem().FieldByName("Title").SetString(title)
	reflect.ValueOf(item).Elem().FieldByName("Description").SetString(description)
	reflect.ValueOf(item).Elem().FieldByName("ImagePath").SetString(path)

	if err := model.DB.Create(item).Error; err != nil {
		return fmt.Errorf("error creating item in the database: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func updateItem(c *fiber.Ctx, modelType interface{}) error {
	// Parse item ID from the request params
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Find the existing item by ID in the database
	item := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()
	if err := model.DB.First(item, id).Error; err != nil {
		// Handle the error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("%s not found", reflect.TypeOf(modelType).Elem().Name())})
	}

	// Parse the item details
	updatedItem := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()
	if err := c.BodyParser(updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Parse the image
	_, updatedPath, err := saveImage(c, false)
	if err != nil {
		return err
	}

	// Update the existing item's details
	itemValue := reflect.ValueOf(item).Elem()
	updatedItemValue := reflect.ValueOf(updatedItem).Elem()

	for i := 0; i < itemValue.NumField(); i++ {
		field := itemValue.Type().Field(i).Name
		if updatedItemValue.FieldByName(field).Interface() != "" {
			itemValue.FieldByName(field).Set(updatedItemValue.FieldByName(field))
		}
	}

	// Update the image path if provided
	if updatedPath != "" {
		itemValue.FieldByName("ImagePath").SetString(updatedPath)
	}

	// Save the updated item back to the database
	model.DB.Save(item)
	return c.Status(fiber.StatusOK).JSON(item)
}


func deleteRecord(c *fiber.Ctx, modelType interface{}) error {
	recordID := c.Params("id")

	// Find the record in the database by ID
	record := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()
	if err := model.DB.First(record, recordID).Error; err != nil {
		// Handle the error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	// Delete the record
	if err := model.DB.Delete(record).Error; err != nil {
		// Handle the error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting record"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Record deleted successfully"})
}

func GetEvents(c *fiber.Ctx) error {
	var events []model.Event
	return getItems(c, &events, "ID")
}


func CreateEvent(c *fiber.Ctx) error {
	return createItem(c, true, &model.Event{})
}

func UpdateEvent(c *fiber.Ctx) error {
	return updateItem(c, &model.Event{})
}

func DeleteEvent(c *fiber.Ctx) error {
	return deleteRecord(c, &model.Event{})
}


func GetNews(c *fiber.Ctx) error {
	var news []model.News
	return getItems(c, &news, "ID")
}

func CreateNews(c *fiber.Ctx) error {
	return createItem(c, true, &model.News{})
}

func UpdateNews(c *fiber.Ctx) error {
	return updateItem(c, &model.News{})
}

func DeleteNews(c *fiber.Ctx) error {
	return deleteRecord(c, &model.News{})
}