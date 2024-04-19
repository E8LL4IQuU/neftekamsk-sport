package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	// Every route here should be added to nginx config at
	// nginx/default.conf, docker container rebuild required afterwards
	// otherwise "404 Not found" will be returned

	// Public routes
	// Auth
	// FIXME: remove "auth" from path
	// app.Post("/api/auth/register", Register)
	app.Post("/api/auth/login", Login)
	// FIXME: what is this
	app.Get("/api/auth/healthcheck", HealthCheck)
	app.Get("/api/auth/user", User)
	app.Post("/api/auth/logout", Logout)

	// Content management
	app.Get("/api/events", GetEvents)
	app.Get("/api/events/:id", GetEventByID)

	app.Get("/api/news", GetNews)
	app.Get("/api/news/:id", GetNewsByID)

	app.Get("/api/signups", GetSignups)
	app.Get("/api/signups/:id", GetSignupByID)

	app.Post("/api/signups", CreateSignup)

	app.Get("/api/photos", GetPhoto)
	app.Get("/api/photos/:id", GetPhotoByID)

	// Authenticated-only routes
	app.Use(NewMiddleware())

	app.Post("/api/events", CreateEvent)
	app.Put("/api/events/:id", UpdateEvent)
	app.Delete("/api/events/:id", DeleteEvent)

	app.Post("/api/news", CreateNews)
	app.Put("/api/news/:id", UpdateNews)
	app.Delete("/api/news/:id", DeleteNews)

	app.Delete("/api/signups/:id", DeleteSignup)

	app.Post("/api/photos", CreatePhoto)
	app.Delete("/api/photos/:id", DeletePhoto)
}
