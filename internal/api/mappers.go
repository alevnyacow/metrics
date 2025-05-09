package api

import (
	"errors"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

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

func (metric Metric) toDomain() domain.Metric {
	getType := func() domain.MetricType {
		if metric.MType == "gauge" {
			return domain.CounterMetricType
		}
		return domain.GaugeMetricType
	}
	getValue := func() string {
		if metric.Delta != nil {
			value, success := domain.CounterRawIntValue(*metric.Delta).ToValue()
			if !success {
				log.Err(errors.New("could not parse value")).Msg("Error on parsing delta")
				return ""
			}
			return value.ToString()
		}
		if metric.Value != nil {
			value, success := domain.GaugeRawFloatValue(*metric.Value).ToValue()
			if !success {
				log.Err(errors.New("could not parse value")).Msg("Error on parsing value")
				return ""
			}
			return value.ToString()
		}
		return ""
	}

	return domain.Metric{
		Name:  metric.ID,
		Value: getValue(),
		Type:  getType(),
	}
}
