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
	store       *session.Store
	AUTH_KEY    string = "authenticated"
	USER_ID     string = "user_id"
	ENVIRONMENT string = os.Getenv("ENVIRONMENT")
)

func InitializeFiber() {
	app := fiber.New(fiber.Config{
		// 100 MB
		BodyLimit: 100 * 1024 * 1024,
		// If raising limit here you will also have to change it in nginx conf for docker containers(neftekamsk-sport/nginx/nginx.conf) and on your host machine, template for it is in /nginx/host.conf.example
	})

	var cookieSecure bool = false
	if ENVIRONMENT == "production" {
		cookieSecure = true
	} else {
		fmt.Println("Warning: running in developer mode, security and auth features disabled")
	}

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   cookieSecure,
		Expiration:     time.Hour * 2880,
		// FIXME: probably has to be anything but None in production
		CookieSameSite: "None",
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	// Auth is inside routes as NewMiddeware()
	InitializeRoutes(app)

	app.Listen(":8000")

	// TODO: Create function that will clear images that are not part of the database anymore
}
