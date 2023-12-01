package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	// Every route should be added to nginx config at
	// nginx/default.conf
	// otherwise "404 Not found" will be returned

	// Auth
	app.Post("/api/register", Register)
	app.Post("/api/login", Login)
	app.Get("/api/user", User)
	app.Post("/api/logout", Logout)

	// Content management
	app.Get("/api/posts", GetEvents)
	app.Post("/api/events", CreateEvent)
}