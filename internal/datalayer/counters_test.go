package datalayer

import "testing"

func TestCounters(t *testing.T) {
	const testCounterName = "testCounterName"
	const testCounterValue CounterMetricValue = 10

	memStorage := NewMemStorage()
	memStorage.AddCounterMetric(testCounterName, testCounterValue)
	val, wasFound := memStorage.GetCounterMetricValue(testCounterName)
	if !wasFound {
		t.Errorf("%s was not found after creating", testCounterName)
	}
	if val != testCounterValue {
		t.Errorf("Expected %s to be %d and got %d", testCounterName, testCounterValue, val)
	}
	memStorage.AddCounterMetric(testCounterName, testCounterValue)
	valAfterAdding, wasFoundAfterAdding := memStorage.GetCounterMetricValue(testCounterName)
	if !wasFoundAfterAdding {
		t.Errorf("%s was not found after modifying", testCounterName)
	}
	if valAfterAdding != testCounterValue*2 {
		t.Errorf("Expected %s to be %d and got %d", testCounterName, testCounterValue*2, val)
	}
}
