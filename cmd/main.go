package main

import (
	"ddd-boilerplate/cmd/config"
	"ddd-boilerplate/infrastructure/fiber"
	"ddd-boilerplate/infrastructure/postgres"
	"ddd-boilerplate/interfaces/http/handler"
	"ddd-boilerplate/interfaces/http/router"
	"ddd-boilerplate/internal/app/service"
	pgInternal "ddd-boilerplate/internal/shared/database/postgres"
	"ddd-boilerplate/pkg/logger"
	"ddd-boilerplate/pkg/migration"
	"log"
)

func main() {
	config := config.NewConfig()
	logger.InitializeLogger()

	psql := postgres.ConnectDB(config.PostgreSQLConfig)
	err := migration.PostgresMigrate()
	if err != nil {
		panic(err)
	}

	sampleRepository := pgInternal.NewSampleRepository(psql)
	sampleService := service.NewSampleService(sampleRepository)

	sampleHandler := handler.NewSampleHandler(sampleService)

	app := fiber.InitFiberApp()

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(config.App.Port))
}
