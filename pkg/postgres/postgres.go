package postgres

import (
	"database/sql"
	"ddd-boilerplate/config"
	"fmt"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDBForMigration(config *config.PostgreSQLConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.HostPrimary,
		config.Username,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to core policy database")
	}

	return db
}

func ConnectDBWithGorm(config *config.PostgreSQLConfig) *gorm.DB {
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
