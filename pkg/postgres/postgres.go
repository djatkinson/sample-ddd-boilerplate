package postgres

import (
	"ddd-boilerplate/config"
	"fmt"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(config *config.PostgreSQLConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.HostPrimary,
		config.Username,
		config.Password,
		config.Name,
		config.Port)

	fmt.Println(dsn)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	return db
}
