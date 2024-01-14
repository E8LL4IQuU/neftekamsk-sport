package routes

import (
	"fmt"
	"os"
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

var environment string

func InitializeFiber() {
	app := fiber.New()

	environment := os.Getenv("ENVIRONMENT")

	var cookieSecure bool = false
	if environment == "production" {
		cookieSecure = true
	} else {
		fmt.Println("Warning: running in developer mode, security features disabled")
	}

	store = session.New(session.Config {
		CookieHTTPOnly: true,
		CookieSecure: cookieSecure,
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
