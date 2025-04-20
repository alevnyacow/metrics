package api

import "github.com/alevnyacow/metrics/internal/domain"

func mTypeFromMetricType(metric domain.Metric) string {
	if metric.IsCounter() {
		return "counter"
	}
	if metric.IsGauge() {
		return "gauge"
	}
	return ""
}

func deltaFromDomainMetric(metric domain.Metric) *int64 {
	if metric.IsCounter() {
		value, parsed := domain.CounterRawValue(metric.Value).ToValue()
		if !parsed {
			return nil
		}
		return (*int64)(&value)
	}
	return nil
}

func valueFromDomainMetric(metric domain.Metric) *float64 {
	if metric.IsGauge() {
		value, parsed := domain.GaugeRawValue(metric.Value).ToValue()
		if !parsed {
			return nil
		}
		return (*float64)(&value)
	}
	return nil
}

func MapDomainMetricToMetricDTO(domainMetric domain.Metric) Metric {
	return Metric{
		ID:    domainMetric.Name,
		MType: mTypeFromMetricType(domainMetric),
		Delta: deltaFromDomainMetric(domainMetric),
		Value: valueFromDomainMetric(domainMetric),
	}
}
