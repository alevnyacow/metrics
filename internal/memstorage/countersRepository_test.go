package memstorage

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func TestCounters(t *testing.T) {
	const testCounterName = "testCounterName"
	const testCounterValue datalayer.CounterValue = 10

	memStorage := NewMemStorage()
	memStorage.AddCounterMetric(testCounterName, testCounterValue)
	val, wasFound := memStorage.GetCounterValue(testCounterName)
	if !wasFound {
		t.Errorf("%s was not found after creating", testCounterName)
	}
	if val != testCounterValue {
		t.Errorf("Expected %s to be %d and got %d", testCounterName, testCounterValue, val)
	}
	memStorage.AddCounterMetric(testCounterName, testCounterValue)
	valAfterAdding, wasFoundAfterAdding := memStorage.GetCounterValue(testCounterName)
	if !wasFoundAfterAdding {
		t.Errorf("%s was not found after modifying", testCounterName)
	}
	if valAfterAdding != testCounterValue*2 {
		t.Errorf("Expected %s to be %d and got %d", testCounterName, testCounterValue*2, val)
	}
}
