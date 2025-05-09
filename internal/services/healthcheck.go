package services

import (
	"database/sql"
	"errors"
)

type HealthcheckService struct {
	db *sql.DB
}

func (service *HealthcheckService) PingDatabase() (success bool, err error) {
	if service.db == nil {
		success = false
		err = errors.New("no database was provided")
		return
	}
	err = service.db.Ping()
	success = err == nil
	return
}
