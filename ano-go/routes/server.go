package routes

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store		*session.Store
	AUTH_KEY	string			= "authenticated"
	USER_ID		string			= "user_id"
)

func InitializeFiber() {
	app := fiber.New()

	store = session.New(session.Config {
		CookieHTTPOnly: true,
		// export to dotenv, true on env=production
		CookieSecure: true,
		Expiration: time.Hour * 2880,
		CookieSameSite:	"None",
	})

	app.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	InitializeRoutes(app)

	app.Listen(":8000")
}
