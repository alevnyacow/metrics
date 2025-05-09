package dbstorage

import (
	"database/sql"
)

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
