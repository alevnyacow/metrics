package domain_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
)

func TestCounterValues(t *testing.T) {
	counterValue := domain.CounterValue(22)
	counterRawValue := domain.CounterRawValue("22")

	counterValueAsString := counterValue.ToString()
	if counterValueAsString != string(counterRawValue) {
		t.Errorf("Wrong string convertation - expected %v, got %v", counterRawValue, counterValueAsString)
	}

	counterValueFromRaw, parsed := counterRawValue.ToValue()
	if !parsed {
		t.Errorf("Could not parse correct value - %v", counterRawValue)
	}
	if counterValueFromRaw != counterValue {
		t.Errorf("Parsed value is incorrect - expected %v, got %v", counterValue, counterValueFromRaw)
	}
}
