package main

import (
	"ddd-boilerplate/cmd/config"
	"ddd-boilerplate/interfaces/http/handler"
	"ddd-boilerplate/interfaces/http/router"
	"ddd-boilerplate/internal/app/service"
	"ddd-boilerplate/internal/shared/database/postgres"
	"log"
)

func main() {
	config.LoadConfig()
	db := config.ConnectDB()

	sampleRepository := postgres.NewSampleRepository(db)

	sampleService := service.NewSampleService(sampleRepository)

	sampleHandler := handler.NewSampleHandler(sampleService)

	app := config.InitFiberApp()

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(":5100"))
}
