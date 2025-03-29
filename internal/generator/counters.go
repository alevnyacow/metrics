package generator

import (
	"fmt"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/datalayer"
	"github.com/alevnyacow/metrics/internal/utils"
)

type Counters struct {
	PollCount datalayer.CounterValue
}

func GenerateCounters() Counters {
	return Counters{
		PollCount: 1,
	}
}

func (counterMetrics *Counters) Links(apiRoot string) (links []string) {
	// Mapping of counter metrics data to their request names.
	counterRequestNamesMapping := map[string]datalayer.CounterValue{
		"PollCount": counterMetrics.PollCount,
	}

	links = make([]string, 0)

	for requestName, metricValue := range counterRequestNamesMapping {
		link := fmt.Sprintf(
			"%s/%s/%s/%s/%s",
			apiRoot,
			api.UpdateLinkRoot,
			api.CounterLinkPath,
			requestName,
			datalayer.CounterValueToString(metricValue),
		)
		links = append(links, link)
	}

	return
}

var _ utils.WithLinks = (*Counters)(nil)
