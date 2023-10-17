package main

import (
	"ddd-boilerplate/interfaces/http/handler"
	"ddd-boilerplate/interfaces/http/router"
	"ddd-boilerplate/internal/app/service"
	"ddd-boilerplate/internal/shared/database/postgres"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {

	sampleRepository := postgres.NewSampleRepository(&gorm.DB{})

	sampleService := service.NewSampleService(sampleRepository)

	sampleHandler := handler.NewSampleHandler(sampleService)

	app := fiber.New()

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(":5100"))
}
