package metrics

import (
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
)

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

		method := c.Request.Method
		route := RouteLabel(c)
		code := strconv.Itoa(c.Writer.Status())

		httpRequestsTotal.WithLabelValues(method, route, code).Inc()
		httpRequestDuration.WithLabelValues(method, route).Observe(time.Since(start).Seconds())
	}
}
