package dbstorage

import (
	"context"
	"database/sql"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

// GaugesRepository is relational database implementation
// of GaugeRepository interface.
type GaugesRepository struct {
	db *sql.DB
}

func (repository *GaugesRepository) PrepareDB(ctx context.Context) error {
	_, err := repository.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS gauges (name TEXT UNIQUE, value DOUBLE PRECISION)`)
	if err != nil {
		log.Err(err).Msg("Error on preparing relational database for gauges repository")
	}
	return err
}

func (repository *GaugesRepository) Set(ctx context.Context, key domain.GaugeName, value domain.GaugeValue) error {
	_, err := repository.db.ExecContext(
		ctx,
		"INSERT INTO gauges (name, value) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET value = $2",
		key,
		value,
	)
	if err != nil {
		log.Err(err).Msg("Error on setting value in relational database for gauges repository")
	}
	return err
}

func (repository *GaugesRepository) Get(ctx context.Context, key domain.GaugeName) (domain.Gauge, error) {
	var gauge domain.Gauge

	row := repository.db.QueryRowContext(ctx, "SELECT name, value FROM gauges WHERE name = $1", key)
	err := row.Scan(&gauge.Name, &gauge.Value)
	if err != nil {
		log.Err(err).Msg("Error on obtaining gauge value from relational database")
	}

	return gauge, err
}

func (repository *GaugesRepository) GetValue(ctx context.Context, key domain.GaugeName) (domain.GaugeValue, error) {
	data, err := repository.Get(ctx, key)
	return data.Value, err
}

func (repository *GaugesRepository) Exists(ctx context.Context, key domain.GaugeName) (bool, error) {
	gauge, err := repository.Get(ctx, key)
	return gauge.Name == key, err
}

func (repository *GaugesRepository) GetAll(ctx context.Context) ([]domain.Gauge, error) {
	gauges := make([]domain.Gauge, 0)
	rows, err := repository.db.QueryContext(ctx, "SELECT name, value FROM gauges")
	if err != nil {
		log.Err(err).Msg("Error on obtaining gauges from relational database")
		return gauges, err
	}
	defer rows.Close()
	for rows.Next() {
		var gauge domain.Gauge
		err := rows.Scan(&gauge.Name, &gauge.Value)
		if err != nil {
			log.Err(err).Msg("Error on parsing metric from relational database")
		}
		gauges = append(gauges, gauge)
	}
	if rows.Err() != nil {
		log.Err(rows.Err()).Msg("Error on obtaining gauges from relational database")
		return gauges, err
	}
	return gauges, err
}
