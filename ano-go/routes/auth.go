package routes

import (
	"strings"

	"github.com/E8LL4IQuU/ano-go/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ParseBody(c *fiber.Ctx) map[string]string {
	var data map[string]string

	c.BodyParser(&data)

	return data
}

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	// We split url at "/" ["", "api", "auth", "login"] to disable auth on login, register routes
	parts := strings.Split(c.Path(), "/")
	// Allow routes with "auth" as second segment of path
	// FIXME: probably "parts[2]" is unsafe and require try
	// sport_backend  | 2024/01/20 08:38:51 stderr: panic: runtime error: index out of range [0] with length 0
	// sport_backend  | 2024/01/20 08:38:51 stderr:
	// sport_backend  | 2024/01/20 08:38:51 stderr: goroutine 10 [running]:
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/E8LL4IQuU/ano-go/routes.CreateEvent(0xc000004600)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /app/routes/content.go:100 +0x207
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/gofiber/fiber/v2.(*App).next(0xc000368000, 0xc000004600)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/gofiber/fiber/v2@v2.50.0/router.go:145 +0x1b2
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/gofiber/fiber/v2.(*Ctx).Next(0xc000288330?)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/gofiber/fiber/v2@v2.50.0/ctx.go:1014 +0x4d
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/gofiber/fiber/v2/middleware/cors.New.func1(0xc000004600)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/gofiber/fiber/v2@v2.50.0/middleware/cors/cors.go:165 +0x3d4
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/gofiber/fiber/v2.(*Ctx).Next(0xc000268ae0?)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/gofiber/fiber/v2@v2.50.0/ctx.go:1011 +0x3d
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/E8LL4IQuU/ano-go/routes.AuthMiddleware(0xc000004600)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /app/routes/auth.go:46 +0x165
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/gofiber/fiber/v2.(*App).next(0xc000368000, 0xc000004600)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/gofiber/fiber/v2@v2.50.0/router.go:145 +0x1b2
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/gofiber/fiber/v2.(*App).handler(0xc000368000, 0x4d1bef?)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/gofiber/fiber/v2@v2.50.0/router.go:172 +0x78
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/valyala/fasthttp.(*Server).serveConn(0xc000320600, {0xa6f338?, 0xc000046410})
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/valyala/fasthttp@v1.50.0/server.go:2359 +0x11d4
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/valyala/fasthttp.(*workerPool).workerFunc(0xc0003068c0, 0xc000268760)
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/valyala/fasthttp@v1.50.0/workerpool.go:224 +0xa4
	// sport_backend  | 2024/01/20 08:38:51 stderr: github.com/valyala/fasthttp.(*workerPool).getCh.func1()
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/valyala/fasthttp@v1.50.0/workerpool.go:196 +0x32
	// sport_backend  | 2024/01/20 08:38:51 stderr: created by github.com/valyala/fasthttp.(*workerPool).getCh in goroutine 1
	// sport_backend  | 2024/01/20 08:38:51 stderr:    /go/pkg/mod/github.com/valyala/fasthttp@v1.50.0/workerpool.go:195 +0x1ab
	if len(parts) > 0 && parts[2] == "auth" {
		return c.Next()
	}

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

func Register(c *fiber.Ctx) error {
	// TODO: check if username, email, password field not empty

	c.Accepts("application/json")
	data := ParseBody(c)

	// Check if email is not taken"
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
	data := ParseBody(c)

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

func HealthCheck(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	auth := sess.Get(AUTH_KEY)
	if auth != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func User(c *fiber.Ctx) error {
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
	var user model.User

	// TODO: check if user exists
	model.DB.Where("id = ?", userId).First(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}
