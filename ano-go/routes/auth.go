package routes

import (
	"strconv"

	"github.com/E8LL4IQuU/ano-go/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func parseBody(c *fiber.Ctx) map[string]string {
	var data map[string]string
	c.BodyParser(&data)

	return data
}

func authMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()
}

func NewMiddleware() fiber.Handler {
	return authMiddleware
}

func Register(c *fiber.Ctx) error {
	// TODO: check if username, email, password field not empty

	c.Accepts("application/json")
	data := parseBody(c)

	// Check if email is not taken
	var user model.User
	model.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email is already taken",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user = model.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	model.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := parseBody(c)

	var user model.User

	model.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	sess, sessErr := store.Get(c)

	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "couldn't login, " + err.Error(),
		})
	}

	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)

	sessErr = sess.Save()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "You have already logged out",
		})
	}

	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "You are logged out",
	})
}

func User(c *fiber.Ctx) error {
	var user model.User

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	userId := sess.Get(USER_ID)
	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	// TODO: check if user exists
	model.DB.Where("id = ?", userId).First(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
	// Get the user ID from the URL parameter
	id := c.Params("id")

	// Convert the ID to an integer
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Find the user by ID and delete it
	var user model.User
	result := model.DB.First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	model.DB.Delete(&user)

	// Return a success response
	return c.SendStatus(fiber.StatusNoContent)
}