package filestorage

import (
	"encoding/json"
	"os"

	"github.com/alevnyacow/metrics/internal/domain"
)

type FileStorage struct {
	path string
}

func (store *FileStorage) Save(metrics []domain.Metric) error {
	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	return os.WriteFile(store.path, data, 0644)
}

func (store *FileStorage) Load() ([]domain.Metric, error) {
	var result []domain.Metric
	res, err := os.ReadFile(store.path)
	if err != nil {
		return result, err
	}
	unmarshallingError := json.Unmarshal(res, &result)
	if unmarshallingError != nil {
		return result, unmarshallingError
	}
	return result, nil
}

func New(path string) *FileStorage {
	return &FileStorage{
		path: path,
	}
}
