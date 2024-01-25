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
	app.Get("/api/events", GetEvents)
	// app.Get("/api/events/:id", GetEventByID)
	app.Post("/api/events", CreateEvent)
	app.Put("/api/events/:id", UpdateEvent)
	app.Delete("/api/events/:id", DeleteEvent)

	app.Get("/api/news", GetNews)
	app.Get("/api/news/:id", GetNewsByID)
	app.Post("/api/news", CreateNews)
	app.Put("/api/news/:id", UpdateNews)
	app.Delete("/api/news/:id", DeleteNews)
}