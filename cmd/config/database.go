package config

import (
	"fmt"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Get("DB_HOST"),
		Get("DB_USER"),
		Get("DB_PASSWORD"),
		Get("DB_NAME"),
		Get("DB_PORT"))

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to core policy database")
	}

	return db
}
