package metrics

import (
	"strconv"
	"time"
)

// CODE_ERROR labels go-pttbbs calls that got no HTTP response
// (connection refused, timeout, ...).
const CODE_ERROR = "error"

// ObserveGoPTTBBS records one HTTP call to go-pttbbs.
//
// route is the route template of the incoming request that triggered the call
// (see RouteLabel). statusCode <= 0 means no response was received.
func ObserveGoPTTBBS(route string, method string, statusCode int, duration time.Duration) {
	code := CODE_ERROR
	if statusCode > 0 {
		code = strconv.Itoa(statusCode)
	}

	gopttbbsRequestsTotal.WithLabelValues(route, method, code).Inc()
	gopttbbsRequestDuration.WithLabelValues(route, method).Observe(duration.Seconds())
}
