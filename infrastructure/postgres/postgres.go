package postgres

import (
	"fmt"
	"os"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
		HostPrimary:   os.Getenv("POSTGRESQL_PRIMARY_HOST"),
		HostSecondary: os.Getenv("POSTGRESQL_SECONDARY_HOST"),
		Port:          os.Getenv("POSTGRESQL_PORT"),
		Username:      os.Getenv("POSTGRESQL_USERNAME"),
		Password:      os.Getenv("POSTGRESQL_PASSWORD"),
		Name:          os.Getenv("POSTGRESQL_NAME"),
	}
}

func ConnectDB(config *PostgreSQLConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.HostPrimary,
		config.Username,
		config.Password,
		config.Name,
		config.Port)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to core policy database")
	}

	return db
}
