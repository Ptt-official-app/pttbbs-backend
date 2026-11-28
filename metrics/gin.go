package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// ROUTE_NO_ROUTE labels requests that matched no registered route
	// (served by NoRoute), so raw paths never become label values.
	ROUTE_NO_ROUTE = "NoRoute"
	// ROUTE_NONE labels calls made outside any HTTP request (cron, queue).
	ROUTE_NONE = "none"
	// METHOD_OTHER labels requests with a non-standard method.
	METHOD_OTHER = "OTHER"
)

var standardMethods = map[string]bool{
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodPatch:   true,
	http.MethodDelete:  true,
	http.MethodConnect: true,
	http.MethodOptions: true,
	http.MethodTrace:   true,
}

// MethodLabel returns method if it is a standard HTTP method, otherwise METHOD_OTHER.
//
// Requests with any method token reach NoRoute, so a client could otherwise
// create unbounded label values (series) just by sending made-up methods.
func MethodLabel(method string) string {
	if !standardMethods[method] {
		return METHOD_OTHER
	}
	return method
}

// RouteLabel returns the matched route template of c (e.g. /api/board/:bid),
// never the raw path, to keep label cardinality bounded.
func RouteLabel(c *gin.Context) string {
	if c == nil {
		return ROUTE_NONE
	}
	route := c.FullPath()
	if route == "" {
		return ROUTE_NO_ROUTE
	}
	return route
}

// GinMiddleware records request count, latency and in-flight requests.
//
// Register it before gin.Recovery() so that recovered panics are counted as 500.
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		httpRequestsInFlight.Inc()
		defer httpRequestsInFlight.Dec()

		c.Next()

		method := MethodLabel(c.Request.Method)
		route := RouteLabel(c)
		code := strconv.Itoa(c.Writer.Status())

		httpRequestsTotal.WithLabelValues(method, route, code).Inc()
		httpRequestDuration.WithLabelValues(method, route).Observe(time.Since(start).Seconds())
	}
}
