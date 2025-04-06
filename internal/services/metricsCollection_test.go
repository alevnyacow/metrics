package services_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
)

func TestBasicCountersCollectionUpdateCase(t *testing.T) {
	service := services.NewMetricsCollectionService()
	if len(service.Counters) != 0 {
		t.Error("Has counters before first update")
	}
	service.UpdateMetrics()
	if len(service.Counters) != 1 {
		t.Error("Collected more counters than should")
	}
	expectedCounterName := domain.CounterName("PollCount")
	expectedCounterValueAfterUpdate := domain.CounterValue(1)
	collectedCounter := service.Counters[0]
	if collectedCounter.Name != expectedCounterName {
		t.Errorf("Collected wrong counter - expected %v, got %v", expectedCounterName, collectedCounter.Name)
	}
	if collectedCounter.Value != expectedCounterValueAfterUpdate {
		t.Errorf("Wrong counter value - expected %v, got %v", expectedCounterValueAfterUpdate, collectedCounter.Value)
	}
}
