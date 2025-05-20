package api

import (
	"fmt"

	"github.com/alevnyacow/metrics/internal/domain"
)

const UpdateLinkRoot = "update"
const ValueLinkRoot = "value"
const pinkLinkRoot = "ping"
const UpdatesLinkRoot = "updates"

const GaugeLinkPath = "gauge"
const CounterLinkPath = "counter"

const typePathParam = "type"
const namePathParam = "name"
const valuePathParam = "value"

func routes() (update string, updateWithJSON string, getMetric string, getAllMetrics string, getByJSON string, ping string, updates string) {
	update = fmt.Sprintf("/%s/{%s}/{%s}/{%s}", UpdateLinkRoot, typePathParam, namePathParam, valuePathParam)
	updateWithJSON = fmt.Sprintf("/%s/", UpdateLinkRoot)
	getMetric = fmt.Sprintf("/%s/{%s}/{%s}", ValueLinkRoot, typePathParam, namePathParam)
	getAllMetrics = "/"
	getByJSON = fmt.Sprintf("/%s/", ValueLinkRoot)
	ping = fmt.Sprintf("/%s", pinkLinkRoot)
	updates = fmt.Sprintf("/%s/", UpdatesLinkRoot)
	return
}

func MetricUpdateByPathRoutes(apiRoot string) (
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

func MetricUpdateByJSONRoute(apiRoot string) string {
	return fmt.Sprintf(
		"%s/%s/",
		apiRoot,
		UpdatesLinkRoot,
	)
}
