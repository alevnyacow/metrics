package memstorage

import "testing"

func TestCounters(t *testing.T) {
	const TEST_COUNTER = "TEST_COUNTER"
	const TEST_VALUE CounterMetricValue = 10

	memStorage := NewMemStorage()
	memStorage.AddCounterMetric(TEST_COUNTER, TEST_VALUE)
	val, wasFound := memStorage.GetCounterMetricValue(TEST_COUNTER)
	if !wasFound {
		t.Errorf("%s was not found after creating", TEST_COUNTER)
	}
	if val != TEST_VALUE {
		t.Errorf("Expected %s to be %d and got %d", TEST_COUNTER, TEST_VALUE, val)
	}
	memStorage.AddCounterMetric(TEST_COUNTER, TEST_VALUE)
	valAfterAdding, wasFoundAfterAdding := memStorage.GetCounterMetricValue(TEST_COUNTER)
	if !wasFoundAfterAdding {
		t.Errorf("%s was not found after modifying", TEST_COUNTER)
	}
	if valAfterAdding != TEST_VALUE*2 {
		t.Errorf("Expected %s to be %d and got %d", TEST_COUNTER, TEST_VALUE*2, val)
	}
}
