package generator

import (
	"fmt"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Agent counter metrics structure contract.
type Counters struct {
	PollCount datalayer.CounterValue
}

// Generates gauges.
func GenerateCounters() Counters {
	return Counters{
		PollCount: 1,
	}
}

// Generates update links for gauge metrics.
func (counterMetrics *Counters) UpdateLinks(apiRoot string) (links []string) {
	// Mapping of counter metrics values to their URL request names.
	counterRequestNamesMapping := map[string]datalayer.CounterValue{
		"PollCount": counterMetrics.PollCount,
	}

	links = make([]string, 0)

	for requestNameInURL, metricValue := range counterRequestNamesMapping {
		link := fmt.Sprintf(
			"%s/%s/%s/%s/%s",
			apiRoot,
			api.UpdateLinkRoot,
			api.CounterLinkPath,
			requestNameInURL,
			datalayer.CounterValueToString(metricValue),
		)
		links = append(links, link)
	}

	return
}

// Static check if Counters implements
// utils.WithLinks interface.
var _ MetricsWithLinks = (*Counters)(nil)
