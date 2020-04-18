package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ServerMetrics struct {
	metrics map[string]prometheus.Collector
}
