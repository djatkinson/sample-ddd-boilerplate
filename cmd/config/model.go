package config

import (
	"ddd-boilerplate/infrastructure/kafka"
	"ddd-boilerplate/infrastructure/oracle"
	"ddd-boilerplate/infrastructure/postgres"
	"ddd-boilerplate/infrastructure/redis"
	"os"
)

type Config struct {
	App *AppConfig

	//DB
	PostgreSQLConfig *postgres.PostgreSQLConfig
	IBORCLConfig     *oracle.OracleDBConfig
	Redis            *redis.RedisConfig
	Kafka            *kafka.KafkaConfig
}

type AppConfig struct {
	Host string
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}
}
