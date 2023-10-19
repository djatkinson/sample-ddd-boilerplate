package http

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/http/handler"
	"ddd-boilerplate/http/router"
	"ddd-boilerplate/internal/app/service"
	pgInternal "ddd-boilerplate/internal/infrastructure/database/postgres"
	"ddd-boilerplate/pkg/fiber"
	"ddd-boilerplate/pkg/metrics"
	"ddd-boilerplate/pkg/postgres"
	"github.com/prometheus/client_golang/prometheus"

	"log"
)

func StartServer(cfg *config.Config) {
	psql := postgres.ConnectDBWithGorm(cfg.PostgreSQLConfig)

	outboundMetrics := metrics.SetupOutboundMetric()
	prometheus.MustRegister(outboundMetrics)

	sampleRepository := pgInternal.NewSampleRepository(psql)
	sampleService := service.NewSampleService(sampleRepository)

	sampleHandler := handler.NewSampleHandler(sampleService)

	app := fiber.InitFiberApp()

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(cfg.App.Port))
}
