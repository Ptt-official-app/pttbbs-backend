package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const namespace = "pttbbs_backend"

// Registry is the registry served on METRICS_HOST.
// A dedicated registry (instead of prometheus.DefaultRegisterer) keeps
// metrics from imported libraries out unless explicitly registered.
var Registry = prometheus.NewRegistry()

var (
	buildInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "build_info",
		Help:      "Build information. Always 1.",
	}, []string{"version", "git_version"})

	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "http",
		Name:      "requests_total",
		Help:      "Number of HTTP requests served, by method, route template and status code.",
	}, []string{"method", "route", "code"})

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "http",
		Name:      "request_duration_seconds",
		Help:      "Latency of HTTP requests served, by method and route template.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"method", "route"})

	httpRequestsInFlight = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: "http",
		Name:      "requests_in_flight",
		Help:      "Number of HTTP requests currently being served.",
	})

	grpcClientRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "grpc_client",
		Name:      "requests_total",
		Help:      "Number of unary gRPC calls to boardd / mand, by full method and gRPC status code.",
	}, []string{"method", "code"})

	grpcClientRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "grpc_client",
		Name:      "request_duration_seconds",
		Help:      "Latency of unary gRPC calls to boardd / mand, by full method.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"method"})

	gopttbbsRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "gopttbbs",
		Name:      "requests_total",
		Help:      "Number of HTTP calls to go-pttbbs, by calling route, method and status code (\"error\" if no response).",
	}, []string{"route", "method", "code"})

	gopttbbsRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "gopttbbs",
		Name:      "request_duration_seconds",
		Help:      "Latency of HTTP calls to go-pttbbs, by calling route and method.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"route", "method"})
)

func init() {
	Registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		buildInfo,
		httpRequestsTotal,
		httpRequestDuration,
		httpRequestsInFlight,
		grpcClientRequestsTotal,
		grpcClientRequestDuration,
		gopttbbsRequestsTotal,
		gopttbbsRequestDuration,
	)
}

// SetBuildInfo exports the build version as pttbbs_backend_build_info.
func SetBuildInfo(version string, gitVersion string) {
	buildInfo.Reset()
	buildInfo.WithLabelValues(version, gitVersion).Set(1)
}
