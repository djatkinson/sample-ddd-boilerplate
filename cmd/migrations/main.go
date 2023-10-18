package main

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/pkg/migration"
	"ddd-boilerplate/pkg/postgres"
)

func main() {
	config := config.NewPostgreSQLConfig()
	dbConn := postgres.ConnectDB(config)
	migration.Execute(dbConn)
}
