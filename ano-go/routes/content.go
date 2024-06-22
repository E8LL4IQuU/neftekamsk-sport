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
	"strings"

	"github.com/E8LL4IQuU/ano-go/model"
	"github.com/gofiber/fiber/v2"
)

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

func hasField(obj interface{}, fieldName string) bool {
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		if typ.Field(i).Name == fieldName {
			return true
		}
	}

	return false
}

func saveImage(c *fiber.Ctx, isImageRequired bool) (string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return "", err
	}
	defer form.RemoveAll()

	files := form.File["image"]

	if isImageRequired {
		if files == nil {
			return "", fmt.Errorf("form file 'image' required")
		}
	} else {
		if files == nil {
			return "", nil
		}
	}

	if len(files) != 1 {
		return "", fmt.Errorf("only one image allowed")
	}

	var path string
	for _, file := range files {
		uploadedFile, err := file.Open()
		if err != nil {
			return "", err
		}
		defer uploadedFile.Close()

		filename, _ := generateSHA256Name(uploadedFile)
		path = filename + getFileExtension(file)
		if err := c.SaveFile(file, "./uploads/"+path); err != nil {
			return "", err
		}
	}

	return path, nil
}

func getItems(c *fiber.Ctx, modelType interface{}, orderByColumn string) error {
	var limit int
	var excludedIDs []uint

	limitParam := c.Query("limit", "10")
	customLimit, err := strconv.Atoi(limitParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit format"})
	}
	limit = customLimit

	excludedIDParam := c.Query("exclude")

	if excludedIDParam != "" {
		ids := strings.Split(excludedIDParam, ",")

		for _, id := range ids {
			parsedID, err := strconv.ParseUint(id, 10, 0) // 0 means to use the bit size of uint
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
			}
			excludedIDs = append(excludedIDs, uint(parsedID))
		}
	}

	modelSliceType := reflect.TypeOf(modelType).Elem()
	modelSlice := reflect.New(modelSliceType).Interface()

	query := model.DB.Order(orderByColumn + " desc").Limit(limit)

	if len(excludedIDs) > 0 {
		query = query.Not("id", excludedIDs)
	}

	if err := query.Find(modelSlice).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(modelSlice)
}

func getItemByID(c *fiber.Ctx, modelType interface{}) error {
	recordID := c.Params("id")

	item := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()

	if err := model.DB.First(item, recordID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func createItem(c *fiber.Ctx, isImageRequired bool, modelType interface{}) error {
	item := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()

	if err := c.BodyParser(item); err != nil {
		return err
	}

	hasImagePath := hasField(item, "ImagePath")

	path, err := saveImage(c, isImageRequired)
	if err != nil {
		return err
	}

	if hasImagePath {
		reflect.ValueOf(item).Elem().FieldByName("ImagePath").SetString(path)
	}

	if err := model.DB.Create(item).Error; err != nil {
		return fmt.Errorf("error creating item in the database: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func updateItem(c *fiber.Ctx, modelType interface{}) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	item := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()
	if err := model.DB.First(item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("%s not found", reflect.TypeOf(modelType).Elem().Name())})
	}

	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	updatedPath, err := saveImage(c, false)
	if err != nil {
		return err
	}

	if updatedPath != "" {
		reflect.ValueOf(item).Elem().FieldByName("ImagePath").SetString(updatedPath)
	}

	model.DB.Save(item)
	return c.Status(fiber.StatusOK).JSON(item)
}

func deleteRecord(c *fiber.Ctx, modelType interface{}) error {
	recordID := c.Params("id")

	record := reflect.New(reflect.TypeOf(modelType).Elem()).Interface()
	if err := model.DB.First(record, recordID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	if err := model.DB.Delete(record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting record"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Record deleted successfully"})
}

func GetEvents(c *fiber.Ctx) error {
	return getItems(c, &[]model.Event{}, "ID")
}

func GetEventByID(c *fiber.Ctx) error {
	return getItemByID(c, &model.Event{})
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
	return getItems(c, &[]model.News{}, "ID")
}

func GetNewsByID(c *fiber.Ctx) error {
	return getItemByID(c, &model.News{})
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

func GetSignups(c *fiber.Ctx) error {
	return getItems(c, &[]model.Signup{}, "ID")
}

func GetSignupByID(c *fiber.Ctx) error {
	return getItemByID(c, &model.Signup{})
}

func CreateSignup(c *fiber.Ctx) error {
	return createItem(c, false, &model.Signup{})
}

func DeleteSignup(c *fiber.Ctx) error {
	return deleteRecord(c, &model.Signup{})
}

func GetPhoto(c *fiber.Ctx) error {
	return getItems(c, &[]model.Photo{}, "ID")
}

func GetPhotoByID(c *fiber.Ctx) error {
	return getItemByID(c, &model.Photo{})
}

func CreatePhoto(c *fiber.Ctx) error {
	return createItem(c, true, &model.Photo{})
}

func DeletePhoto(c *fiber.Ctx) error {
	return deleteRecord(c, &model.Photo{})
}

func GetAthletes(c *fiber.Ctx) error {
	return getItems(c, &[]model.Athlete{}, "ID")
}

func GetAthleteByID(c *fiber.Ctx) error {
	return getItemByID(c, &model.Athlete{})
}

func CreateAthlete(c *fiber.Ctx) error {
	return createItem(c, true, &model.Athlete{})
}

func UpdateAthlete(c *fiber.Ctx) error {
	return updateItem(c, &model.Athlete{})
}

func DeleteAthlete(c *fiber.Ctx) error {
	return deleteRecord(c, &model.Athlete{})
}
