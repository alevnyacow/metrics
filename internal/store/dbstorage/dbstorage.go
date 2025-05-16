package dbstorage

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

func InitDatabase(connectionString string) (db *sql.DB, closeConnection func()) {
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = database
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
