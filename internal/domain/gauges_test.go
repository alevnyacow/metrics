package domain_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
)

func TestGaugeValues(t *testing.T) {
	gaugeValue := domain.GaugeValue(22.5)
	gaugeRawValue := domain.GaugeRawValue("22.5")

	gaugeValueAsString := gaugeValue.ToString()
	if gaugeValueAsString != string(gaugeRawValue) {
		t.Errorf("Wrong string convertation - expected %v, got %v", gaugeRawValue, gaugeValueAsString)
	}

	gaugeValueFromRaw, parsed := gaugeRawValue.ToValue()
	if !parsed {
		t.Errorf("Could not parse correct value - %v", gaugeRawValue)
	}
	if gaugeValueFromRaw != gaugeValue {
		t.Errorf("Parsed value is incorrect - expected %v, got %v", gaugeValue, gaugeValueFromRaw)
	}
}
