package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App *AppConfig

	//DB
	PostgreSQLConfig *PostgreSQLConfig
	IBORCLConfig     *OracleDBConfig
	Redis            *RedisConfig
	Kafka            *KafkaConfig
}

type AppConfig struct {
	Host string
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Host: viper.GetString("APP_HOST"),
		Port: viper.GetString("APP_PORT"),
	}
}

type PostgreSQLConfig struct {
	HostPrimary   string
	HostSecondary string
	Port          string
	Name          string
	Username      string
	Password      string
}

func NewPostgreSQLConfig() *PostgreSQLConfig {
	return &PostgreSQLConfig{
		HostPrimary:   viper.GetString("POSTGRESQL_PRIMARY_HOST"),
		HostSecondary: viper.GetString("POSTGRESQL_SECONDARY_HOST"),
		Port:          viper.GetString("POSTGRESQL_PORT"),
		Username:      viper.GetString("POSTGRESQL_USERNAME"),
		Password:      viper.GetString("POSTGRESQL_PASSWORD"),
		Name:          viper.GetString("POSTGRESQL_DATABASE_NAME"),
	}
}

type OracleDBConfig struct {
	HostPrimary   string
	HostSecondary string
	Port          int
	Name          string
	Username      string
	Password      string
}

type RedisConfig struct {
	Host string
	DB   string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host: viper.GetString("REDIS_HOST"),
		DB:   viper.GetString("REDIS_DB"),
	}
}

type KafkaConfig struct {
	BrokerURL string
	Partition int
}
