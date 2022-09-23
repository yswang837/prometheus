package client

import "github.com/yswang837/prometheus/metrics"

var HelloCounter *helloCounter

type helloCounter struct {
	metricsHelloCounter *metrics.Counter
}

type MetricsHelloTagsCounter struct {
	Name string
}

func newHelloCounter() *helloCounter {
	count := &helloCounter{
		metricsHelloCounter: metrics.NewCounter("helloCounter", []string{}),
	}
	return count
}

func (h *helloCounter) IncMetricsHelloCounter(tags *MetricsHelloTagsCounter) {
	h.metricsHelloCounter.Attributes(tags.Name).Inc()
}

func init() {
	HelloCounter = newHelloCounter()
}
