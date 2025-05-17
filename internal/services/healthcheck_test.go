package services_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/services"
)

func TestHealthcheckWithoutDatabase(t *testing.T) {
	service := services.NewHealtheckService(nil)
	success, error := service.PingDatabase()
	if success {
		t.Error("Success in error case")
	}
	if error == nil {
		t.Error("No errors were returned")
	}
}
