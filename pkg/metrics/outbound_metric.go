package metrics

import "github.com/prometheus/client_golang/prometheus"

func SetupOutboundMetric() *prometheus.CounterVec {
	outboundCall := prometheus.CounterOpts{
		Name: "http_request_outbound_callbacks",
		Help: "Get Response information from outbound http request",
	}
	outboundCallLabelName := []string{
		"host",
		"path",
		"response_code",
	}

	metricOutbound := prometheus.NewCounterVec(outboundCall, outboundCallLabelName)
	return metricOutbound
}
