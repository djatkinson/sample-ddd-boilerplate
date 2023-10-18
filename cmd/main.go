package main

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/interfaces/http/handler"
	"ddd-boilerplate/interfaces/http/router"
	"ddd-boilerplate/internal/app/service"
	pgInternal "ddd-boilerplate/internal/shared/database/postgres"
	"ddd-boilerplate/pkg/fiber"
	"ddd-boilerplate/pkg/logger"
	"ddd-boilerplate/pkg/postgres"
	"log"
)

func main() {
	config := config.NewConfig()
	logger.InitializeLogger()

	psql := postgres.ConnectDB(config.PostgreSQLConfig)
	sampleRepository := pgInternal.NewSampleRepository(psql)

	sampleService := service.NewSampleService(sampleRepository)
	sampleHandler := handler.NewSampleHandler(sampleService)

	app := fiber.InitFiberApp()

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(config.App.Port))
}
