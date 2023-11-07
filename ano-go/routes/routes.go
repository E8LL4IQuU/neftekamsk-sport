package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {

	app.Post("/api/register", Register)
	app.Post("/api/login", Login)
	app.Get("/api/user", User)
	app.Post("/api/logout", Logout)
}