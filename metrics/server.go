package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const METRICS_R = "/metrics"

// Handler serves Registry in the Prometheus exposition format.
func Handler() http.Handler {
	return promhttp.HandlerFor(Registry, promhttp.HandlerOpts{Registry: Registry})
}

// NewServer returns a server exposing METRICS_R on addr.
//
// It is meant to listen on an internal-only address, separate from the public API.
func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle(METRICS_R, Handler())

	return &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
	}
}
