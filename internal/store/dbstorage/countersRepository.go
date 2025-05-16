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

func (repository *CountersRepository) PrepareDB(ctx context.Context) error {
	_, err := repository.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS counters (name TEXT UNIQUE, value BIGINT)`)
	if err != nil {
		log.Err(err).Msg("Error on preparing relational database for counters repository")
	}
	return err
}

func (repository *CountersRepository) Set(ctx context.Context, key domain.CounterName, value domain.CounterValue) error {
	_, err := repository.db.ExecContext(
		ctx,
		"INSERT INTO counters (name, value) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET value = $2",
		key,
		value,
	)
	return err
}

func (repository *CountersRepository) Get(ctx context.Context, key domain.CounterName) (domain.Counter, error) {
	var counter domain.Counter

	row := repository.db.QueryRowContext(ctx, "SELECT name, value FROM counters WHERE name = $1", key)
	err := row.Scan(&counter.Name, &counter.Value)

	return counter, err
}

func (repository *CountersRepository) GetValue(ctx context.Context, key domain.CounterName) (domain.CounterValue, error) {
	data, error := repository.Get(ctx, key)
	return data.Value, error
}

func (repository *CountersRepository) Exists(ctx context.Context, key domain.CounterName) (bool, error) {
	counter, error := repository.Get(ctx, key)
	return counter.Name == key, error
}

func (repository *CountersRepository) GetAll(ctx context.Context) ([]domain.Counter, error) {
	counters := make([]domain.Counter, 0)
	rows, err := repository.db.QueryContext(ctx, "SELECT name, value FROM counters")
	if err != nil {
		log.Err(err).Msg("Error on obtaining counters from relational database")
		return counters, err
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
	if rows.Err() != nil {
		log.Err(rows.Err()).Msg("Error on obtaining counters from relational database")
		return counters, rows.Err()
	}
	return counters, nil
}
