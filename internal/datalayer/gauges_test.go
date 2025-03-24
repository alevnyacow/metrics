package datalayer

import "testing"

func TestGauges(t *testing.T) {
	const TEST_GAUGE = "TEST_GAUGE"
	const TEST_VALUE GaugeMetricValue = 10

	memStorage := NewMemStorage()
	memStorage.SetGaugeMetric(TEST_GAUGE, TEST_VALUE)
	val, wasFound := memStorage.GetGaugeMetricValue(TEST_GAUGE)
	if !wasFound {
		t.Errorf("%s was not found after creating", TEST_GAUGE)
	}
	if val != TEST_VALUE {
		t.Errorf("Expected %s to be %f and got %f", TEST_GAUGE, TEST_VALUE, val)
	}
	memStorage.SetGaugeMetric(TEST_GAUGE, TEST_VALUE+1)
	valAfterAdding, wasFoundAfterAdding := memStorage.GetGaugeMetricValue(TEST_GAUGE)
	if !wasFoundAfterAdding {
		t.Errorf("%s was not found after modifying", TEST_GAUGE)
	}
	if valAfterAdding != TEST_VALUE+1 {
		t.Errorf("Expected %s to be %f and got %f", TEST_GAUGE, TEST_VALUE+1, valAfterAdding)
	}
}
