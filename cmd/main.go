package main

import (
	"ddd-boilerplate/cmd/config"
	"ddd-boilerplate/infrastructure/fiber"
	"ddd-boilerplate/infrastructure/postgres"
	"ddd-boilerplate/interface/http/handler"
	"ddd-boilerplate/interface/http/router"
	"ddd-boilerplate/internal/app/service"
	pgInternal "ddd-boilerplate/internal/shared/database/postgres"
	"ddd-boilerplate/pkg/logger"
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
