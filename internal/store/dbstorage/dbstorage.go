package dbstorage

import (
	"database/sql"
	"fmt"

	"github.com/alevnyacow/metrics/internal/retries"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func InitDatabase(connectionString string, migrationsPath string) (db *sql.DB, closeConnection func()) {
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = database
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Err(err).Msg("Error creating driver")
	}
	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres",
		driver,
	)

	if err != nil {
		log.Err(err).Msg("Error on migration initialization")
	}

	if errOnUp := retries.WithRetries(migration.Up); errOnUp != nil {
		log.Err(errOnUp).Msg("Error on applying migrations")
		errOnDown := retries.WithRetries(migration.Down)
		if errOnDown != nil {
			log.Err(errOnDown).Msg("Error on dropping migrations")
		}
	}

	closeConnection = func() {
		if db != nil {
			err := db.Close()
			if err != nil {
				log.Err(err).Msg("Error on closing database")
			}
		}
	}
	return
}

func NewCountersRepository(db *sql.DB) *CountersRepository {
	return &CountersRepository{
		db: db,
	}
}

func NewGaugesRepository(db *sql.DB) *GaugesRepository {
	return &GaugesRepository{
		db: db,
	}
}
