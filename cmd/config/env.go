package config

import (
	"ddd-boilerplate/infrastructure/postgres"
	"ddd-boilerplate/pkg/logger"

	"github.com/joho/godotenv"
)

func NewConfig() *Config {
	LoadConfigFromFile()
	return &Config{
		App:              NewAppConfig(),
		PostgreSQLConfig: postgres.NewPostgreSQLConfig(),
	}
}

func LoadConfigFromFile() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}
