package config

import (
	"ddd-boilerplate/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
