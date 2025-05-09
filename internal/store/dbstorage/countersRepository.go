package dbstorage

import (
	"context"
	"database/sql"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

// CounterRepository is relational database implementation
// of CounterRepository interface.
type CountersRepository struct {
	db *sql.DB
}

func (repository *CountersRepository) PrepareDB(ctx context.Context) {
	_, err := repository.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS counters (name TEXT UNIQUE, value INTEGER)`)
	if err != nil {
		log.Err(err).Msg("Error on preparing relational database for counters repository")
	}
}

func (repository *CountersRepository) Set(ctx context.Context, key domain.CounterName, value domain.CounterValue) {
	repository.db.ExecContext(
		ctx,
		"INSERT INTO counters (name, value) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET value = $2",
		key,
		value,
	)
}

func (repository *CountersRepository) Get(ctx context.Context, key domain.CounterName) domain.Counter {
	var counter domain.Counter

	row := repository.db.QueryRowContext(ctx, "SELECT name, value FROM counters WHERE name = $1", key)
	err := row.Scan(&counter.Name, &counter.Value)
	if err != nil {
		log.Err(err).Msg("Error on obtaining counter value from relational database")
	}

	return counter
}

func (repository *CountersRepository) GetValue(ctx context.Context, key domain.CounterName) domain.CounterValue {
	data := repository.Get(ctx, key)
	return data.Value
}

func (repository *CountersRepository) Exists(ctx context.Context, key domain.CounterName) bool {
	counter := repository.Get(ctx, key)
	return counter.Name == key
}

func (repository *CountersRepository) GetAll(ctx context.Context) []domain.Counter {
	counters := make([]domain.Counter, 0)
	rows, err := repository.db.QueryContext(ctx, "SELECT name, value FROM counters")
	if err != nil {
		log.Err(err).Msg("Error on obtaining metrics data from relational database")
		return counters
	}
	defer rows.Close()
	for rows.Next() {
		var counter domain.Counter
		err := rows.Scan(&counter.Name, &counter.Value)
		if err != nil {
			log.Err(err).Msg("Error on parsing metric from relational database")
		}
		counters = append(counters, counter)
	}
	return counters
}
