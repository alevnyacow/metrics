package memstorage

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func TestGauges(t *testing.T) {
	const testGaugeName = "testGaugeName"
	const testGaugeValue datalayer.GaugeValue = 10

	memStorage := NewMemStorage()
	memStorage.SetGaugeMetric(testGaugeName, testGaugeValue)
	val, wasFound := memStorage.GetGaugeValue(testGaugeName)
	if !wasFound {
		t.Errorf("%s was not found after creating", testGaugeName)
	}
	if val != testGaugeValue {
		t.Errorf("Expected %s to be %f and got %f", testGaugeName, testGaugeValue, val)
	}
	memStorage.SetGaugeMetric(testGaugeName, testGaugeValue+1)
	valAfterAdding, wasFoundAfterAdding := memStorage.GetGaugeValue(testGaugeName)
	if !wasFoundAfterAdding {
		t.Errorf("%s was not found after modifying", testGaugeName)
	}
	if valAfterAdding != testGaugeValue+1 {
		t.Errorf("Expected %s to be %f and got %f", testGaugeName, testGaugeValue+1, valAfterAdding)
	}
}
