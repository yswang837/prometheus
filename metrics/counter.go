package metrics

import "github.com/prometheus/client_golang/prometheus"

type Counter struct {
	promCounterVec *prometheus.CounterVec
}

func NewCounter(name string, tags []string) *Counter {
	tags = append([]string{}, tags...)
	opts := prometheus.CounterOpts{Name: name}
	vector := prometheus.NewCounterVec(opts, tags)
	c := &Counter{vector}
	prometheus.MustRegister(vector)
	return c
}
func (c *Counter) Attributes(attrs ...string) prometheus.Counter {
	attrs = append([]string{}, attrs...)
	return c.promCounterVec.WithLabelValues(attrs...)
}
