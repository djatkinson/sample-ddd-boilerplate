package migration

import (
	"database/sql"
	"ddd-boilerplate/pkg/logger"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func PostgresMigrate() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRESQL_PRIMARY_HOST"),
		os.Getenv("POSTGRESQL_USERNAME"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_DATABASE_NAME"),
		os.Getenv("POSTGRESQL_PORT"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Logger.Error(err.Error())
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Logger.Error(err.Error())
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres", driver)
	if err != nil {
		logger.Logger.Error(err.Error())
		return err
	}

	startVersion, _, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		logger.Logger.Error(err.Error())
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked {
		log.Error("Migration is dirty, forcing rollback and retrying")
		endVersion, _, _ := m.Version()
		fmt.Println("startversion", startVersion, "endversion", endVersion)
		m.Steps(0 - 4)
	}

	return nil
}

func Rollback(startVersion uint) {

}
