package http

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/http/handler"
	"ddd-boilerplate/http/router"
	pgInternal "ddd-boilerplate/internal/app/database/postgres"
	"ddd-boilerplate/internal/app/service"
	"ddd-boilerplate/pkg/fiber"
	"ddd-boilerplate/pkg/postgres"

	"log"
)

func StartServer(cfg *config.Config) {
	psql := postgres.ConnectDBWithGorm(cfg.PostgreSQLConfig)

	sampleRepository := pgInternal.NewSampleRepository(psql)
	sampleService := service.NewSampleService(sampleRepository)

	sampleHandler := handler.NewSampleHandler(sampleService)

	app := fiber.InitFiberApp()

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(cfg.App.Port))
}