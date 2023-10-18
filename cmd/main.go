package main

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/http"
	"ddd-boilerplate/pkg/logger"
)

func main() {
	cfg := config.NewConfig()
	logger.InitializeLogger()

	http.StartServer(cfg)
}
