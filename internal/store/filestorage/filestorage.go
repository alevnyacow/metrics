package filestorage

import (
	"encoding/json"
	"os"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

type FileStorage struct {
	path string
}

func (store *FileStorage) Save(metrics []domain.Metric) error {
	data, err := json.Marshal(metrics)
	if err != nil {
		log.Err(err).Msg("Saving in file storage")
		return err
	}

	return os.WriteFile(store.path, data, 0644)
}

func (store *FileStorage) Load() ([]domain.Metric, error) {
	var result []domain.Metric
	res, err := os.ReadFile(store.path)
	if err != nil {
		log.Err(err).Msg("Reading file of file storage")
		return result, err
	}
	unmarshallingError := json.Unmarshal(res, &result)
	if unmarshallingError != nil {
		log.Err(unmarshallingError).Msg("Unmarshaling data from file storage")
		return result, unmarshallingError
	}
	return result, nil
}

func New(path string) *FileStorage {
	return &FileStorage{
		path: path,
	}
}
