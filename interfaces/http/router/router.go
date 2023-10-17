package router

import (
	"ddd-boilerplate/interfaces/http/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, sampleHandler handler.SampleHandler) {
	sampleRouter := app.Group("/sample")
	sampleRouter.Get("/:id", sampleHandler.GetSampleByID)
}
