package generator

// Contract of metrics collection which can
// generate links based on provided API root
// and their internal state.
type MetricsWithLinks interface {
	// Links for metrics updating.
	UpdateLinks(apiRoot string) (links []string)
}
