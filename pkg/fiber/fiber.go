package fiber

import (
	"ddd-boilerplate/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.elastic.co/apm/module/apmfiber/v2"
)

func InitFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(apmfiber.Middleware())
	app.Use(middleware.Logger())
	return app
}
