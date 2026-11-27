package metrics

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(GinMiddleware(), gin.Recovery())

	router.GET("/test/board/:bid", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/test/panic", func(_ *gin.Context) {
		panic("test panic")
	})
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "index.html")
	})

	return router
}

func serve(router *gin.Engine, path string) int {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	router.ServeHTTP(w, req)
	return w.Code
}

func TestGinMiddleware(t *testing.T) {
	router := newTestRouter()

	tests := []struct {
		name     string
		paths    []string
		route    string
		code     string
		expected float64
	}{
		{
			name:     "labels by route template, not raw path",
			paths:    []string{"/test/board/1", "/test/board/2"},
			route:    "/test/board/:bid",
			code:     "200",
			expected: 2,
		},
		{
			name:     "recovered panic counts as 500",
			paths:    []string{"/test/panic"},
			route:    "/test/panic",
			code:     "500",
			expected: 1,
		},
		{
			name:     "unmatched paths share one label",
			paths:    []string{"/some/spa/path", "/another/path"},
			route:    ROUTE_NO_ROUTE,
			code:     "200",
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := testutil.ToFloat64(httpRequestsTotal.WithLabelValues(http.MethodGet, tt.route, tt.code))
			for _, each := range tt.paths {
				serve(router, each)
			}
			after := testutil.ToFloat64(httpRequestsTotal.WithLabelValues(http.MethodGet, tt.route, tt.code))
			assert.Equal(t, tt.expected, after-before)
		})
	}

	assert.Equal(t, float64(0), testutil.ToFloat64(httpRequestsInFlight))
}

func TestRouteLabel(t *testing.T) {
	assert.Equal(t, ROUTE_NONE, RouteLabel(nil))
}

func TestUnaryClientInterceptor(t *testing.T) {
	interceptor := UnaryClientInterceptor()
	const method = "/test.Service/Test"

	tests := []struct {
		name string
		err  error
		code string
	}{
		{name: "ok", err: nil, code: codes.OK.String()},
		{name: "unavailable", err: status.Error(codes.Unavailable, "down"), code: codes.Unavailable.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoker := func(_ context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
				return tt.err
			}

			before := testutil.ToFloat64(grpcClientRequestsTotal.WithLabelValues(method, tt.code))
			err := interceptor(context.Background(), method, nil, nil, nil, invoker)
			assert.Equal(t, tt.err, err)
			after := testutil.ToFloat64(grpcClientRequestsTotal.WithLabelValues(method, tt.code))
			assert.Equal(t, float64(1), after-before)
		})
	}
}

func TestObserveGoPTTBBS(t *testing.T) {
	const route = "/test/gopttbbs"

	tests := []struct {
		name       string
		statusCode int
		code       string
	}{
		{name: "response", statusCode: 403, code: "403"},
		{name: "no response", statusCode: 0, code: CODE_ERROR},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := testutil.ToFloat64(gopttbbsRequestsTotal.WithLabelValues(route, http.MethodPost, tt.code))
			ObserveGoPTTBBS(route, http.MethodPost, tt.statusCode, time.Millisecond)
			after := testutil.ToFloat64(gopttbbsRequestsTotal.WithLabelValues(route, http.MethodPost, tt.code))
			assert.Equal(t, float64(1), after-before)
		})
	}
}

func TestHandler(t *testing.T) {
	SetBuildInfo("v-test", "abc123")
	serve(newTestRouter(), "/test/board/1")

	srv := httptest.NewServer(NewServer("").Handler)
	defer srv.Close()

	resp, err := http.Get(srv.URL + METRICS_R)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	for _, each := range []string{
		`pttbbs_backend_build_info{git_version="abc123",version="v-test"} 1`,
		`pttbbs_backend_http_requests_total{code="200",method="GET",route="/test/board/:bid"}`,
		"pttbbs_backend_http_request_duration_seconds_bucket",
		"go_goroutines",
	} {
		assert.Contains(t, string(body), each)
	}
}
