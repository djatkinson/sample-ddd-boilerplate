package migration

import (
	"database/sql"
	"ddd-boilerplate/pkg/logger"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func PostgresMigrate() error {
	log := logger.Logger
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRESQL_PRIMARY_HOST"),
		os.Getenv("POSTGRESQL_USERNAME"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_DATABASE_NAME"),
		os.Getenv("POSTGRESQL_PORT"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres", driver)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	startVersion, _, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		log.Error(err.Error())
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked {
		log.Error("Migration is dirty, forcing rollback and retrying")
		endVersion, _, err := m.Version()
		if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
			log.Error(err.Error())
			return err
		}

		m.Force(int(endVersion) - 1)
		m.Steps((int(startVersion) + 1) - int(endVersion))
		m.Force(int(startVersion))
	}

	return nil
}
