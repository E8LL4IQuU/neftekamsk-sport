package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	// Every route should be added to nginx config at
	// nginx/default.conf
	// otherwise "404 Not found" will be returned

	// Auth
	// /api/auth routes are allowed to access without authorization
	app.Post("/api/auth/register", Register)
	app.Post("/api/auth/login", Login)
	app.Get("/api/auth/user", User)
	app.Post("/api/auth/logout", Logout)

	// Content management
	app.Get("/api/posts", GetEvents)
	app.Post("/api/events", CreateEvent)
	// TODO: rewrite for rest api, app.Get("/api/events/:id")
	app.Get("/api/events", GetEvents)
	app.Delete("/api/events/:id", DeleteEvent)
}