package api

import (
	"fmt"

	"github.com/alevnyacow/metrics/internal/domain"
)

const UpdateLinkRoot = "update"
const ValueLinkRoot = "value"

const GaugeLinkPath = "gauge"
const CounterLinkPath = "counter"

const typePathParam = "type"
const namePathParam = "name"
const valuePathParam = "value"

func routes() (update string, getMetric string, getAllMetrics string) {
	update = fmt.Sprintf("/%s/{%s}/{%s}/{%s}", UpdateLinkRoot, typePathParam, namePathParam, valuePathParam)
	getMetric = fmt.Sprintf("/%s/{%s}/{%s}", ValueLinkRoot, typePathParam, namePathParam)
	getAllMetrics = "/"
	return
}

func MetricUpdateRoutes(apiRoot string) (
	forCounter func(counter domain.Counter) string,
	forGauge func(gauge domain.Gauge) string,
) {
	forCounter = func(counter domain.Counter) string {
		return fmt.Sprintf(
			"%s/%s/%s/%s/%s",
			apiRoot,
			UpdateLinkRoot,
			CounterLinkPath,
			counter.Name,
			counter.Value.ToString(),
		)
	}
	forGauge = func(gauge domain.Gauge) string {
		return fmt.Sprintf(
			"%s/%s/%s/%s/%s",
			apiRoot,
			UpdateLinkRoot,
			GaugeLinkPath,
			gauge.Name,
			gauge.Value.ToString(),
		)
	}
	return
}
